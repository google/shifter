// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package openshift

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	k8sjson "k8s.io/apimachinery/pkg/runtime/serializer/json"
	types "k8s.io/apimachinery/pkg/types"
	restclientcmdapi "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	appsv1 "github.com/openshift/api/apps/v1"
	authorizationv1 "github.com/openshift/api/authorization/v1"
	buildv1 "github.com/openshift/api/build/v1"
	imagev1 "github.com/openshift/api/image/v1"
	networkv1 "github.com/openshift/api/network/v1"
	oauthv1 "github.com/openshift/api/oauth/v1"
	projectv1 "github.com/openshift/api/project/v1"
	quotav1 "github.com/openshift/api/quota/v1"
	routev1 "github.com/openshift/api/route/v1"
	securityv1 "github.com/openshift/api/security/v1"
	templatev1 "github.com/openshift/api/template/v1"
	userv1 "github.com/openshift/api/user/v1"

	"shifter/generator"
	"shifter/processor"

	"k8s.io/client-go/kubernetes/scheme"

	"shifter/lib"
	//"reflect"
	//"fmt"
	"shifter/ops"
)

type Openshift struct {
	Endpoint  string
	AuthToken string
	Username  string
	Password  string
}

type ResourceList struct {
	Namespace  string
	Kind       string
	APIVersion string
	Name       string
	UID        types.UID
	Payload    bytes.Buffer
	PayloadStr string
	Error      error
}

func (c *Openshift) clusterClient() *restclientcmdapi.Config {
	config := clientcmdapi.NewConfig()
	config.Clusters["cluster"] = &clientcmdapi.Cluster{
		InsecureSkipTLSVerify: true,
		Server:                c.Endpoint,
	}

	if len(c.AuthToken) <= 5 {
		lib.CLog("error", "Invalid OpenShift Token", nil)
		os.Exit(1) // TODO - RETURN ERROR
	}

	config.AuthInfos["cluster-auth"] = &clientcmdapi.AuthInfo{
		Token: c.AuthToken,
	}

	config.Contexts["ctx"] = &clientcmdapi.Context{
		Cluster:  "cluster",
		AuthInfo: "cluster-auth",
	}

	config.CurrentContext = "ctx"
	clusterConfig := clientcmd.NewNonInteractiveClientConfig(*config, "ctx", &clientcmd.ConfigOverrides{}, nil)

	cl, err := clusterConfig.ClientConfig()
	if err != nil {
		lib.CLog("error", "Configuring OpenShift Cluster Client", err)
	}
	return cl
}

func (c *Openshift) ExportNSResources(namespace string, outputPath string) error {
	var (
		resourcelist []ResourceList
		err          error
	)

	if namespace != "" {
		resourcelist, err = c.getResources(namespace, true)
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		projects, err := c.GetAllProjects()
		if err != nil {
			lib.CLog("error", "Unable to get all projects", err)
			return err
		}
		for _, p := range projects.Items {
			rl, err := c.getResources(p.ObjectMeta.Name, true)
			if err != nil {
				log.Println(err)
				return err
			}
			for _, v := range rl {
				resourcelist = append(resourcelist, v)
			}
		}
	}
	for _, res := range resourcelist {
		// TODO - Hard Coded Values (YAML - LOCAL)
		lib.CLog("info", "Exporting object: "+res.Namespace+"\\"+res.Name+" of kind "+res.Kind)
		fileObj := &ops.FileObject{
			StorageType:   "local",
			Path:          (outputPath + "/" + res.Namespace + "/" + res.Name),
			Ext:           "yaml",
			Content:       res.Payload,
			ContentLength: res.Payload.Len(),
		}
		err := fileObj.WriteFile()
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Openshift) ConvertNSResources(namespace string, flags map[string]string, outputPath string) error {
	var (
		resourcelist []ResourceList
		err          error
	)

	if namespace != "" {
		resourcelist, err = c.getResources(namespace, false)
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		projects, err := c.GetAllProjects()
		if err != nil {
			lib.CLog("error", "Unable to get projects", err)
			return err
		}
		for _, p := range projects.Items {
			rl, err := c.getResources(p.ObjectMeta.Name, false)
			if err != nil {
				log.Println(err)
				return err
			}
			for _, v := range rl {
				resourcelist = append(resourcelist, v)
			}
		}
	}

	for _, res := range resourcelist {
		lib.CLog("info", "Converting object: "+res.Namespace+"\\"+res.Name+" of kind "+res.Kind)

		obj, err := processor.Processor(res.Payload.Bytes(), res.Kind, flags)
		if err != nil {
			lib.CLog("error", "Cannot create shifter processor.", err)
			return err
		}

		convertedObject, err := generator.NewGenerator("yaml", res.Name, obj)
		if err != nil {
			lib.CLog("error", "Cannot create shifter generator.", err)
			return err
		}
		// TODO - Hard Coded Values (YAML - LOCAL)
		for _, conObj := range convertedObject {
			fileObj := &ops.FileObject{
				StorageType:   "local",
				Path:          (outputPath + "/" + res.Namespace + "/" + conObj.Name),
				Ext:           "yaml",
				Content:       conObj.Payload,
				ContentLength: conObj.Payload.Len(),
			}
			err := fileObj.WriteFile()
			if err != nil {
				return err
			}
		}
	}
	return err
}

func (c *Openshift) ListNSResources(csvoutput bool, namespace string) error {
	var (
		resourcelist []ResourceList
		err          error
	)

	if namespace != "" {
		resourcelist, err = c.getResources(namespace, false)
		if err != nil {
			lib.CLog("error", "Getting resources from namespace "+namespace, err)
		}
	} else {
		projects, err := c.GetAllProjects()
		if err != nil {
			lib.CLog("error", "Unable to access OpenShift Project", err)
			return err
		}
		for _, p := range projects.Items {
			rl, err := c.getResources(p.ObjectMeta.Name, false)
			if err != nil {
				lib.CLog("error", "Getting project "+p.ObjectMeta.Name, err)
				return err
			}
			for _, v := range rl {
				resourcelist = append(resourcelist, v)
			}
		}
	}

	t := table.NewWriter()
	t.SetStyle(table.StyleColoredBright)
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Project", "Kind", "Name", "Payload", "Error"})

	for _, x := range resourcelist {
		t.AppendRow(table.Row{x.Namespace, x.Kind, x.Name, len(x.Payload.String()), x.Error})
	}
	if csvoutput == false {
		t.Render()
	} else {
		t.RenderCSV()
	}

	return nil
}

func (c *Openshift) GetResources(namespace string, yaml bool, kind string, name string, uid types.UID) ([]ResourceList, error) {

	kind = strings.ToLower(kind)
	namespace = strings.ToLower(namespace)
	name = strings.ToLower(name)

	var resources []ResourceList
	rl, err := c.getResources(namespace, yaml)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, v := range rl {
		if uid == "" || uid == v.UID {
			if kind == "" || kind == strings.ToLower(v.Kind) {
				if name == "" || name == strings.ToLower(v.Name) {
					resources = append(resources, v)
				}
			}
		}
	}

	return resources, err
}

func (c *Openshift) getResources(namespace string, yaml bool) ([]ResourceList, error) {
	var (
		resourcelist []ResourceList
		err          error
	)

	// Add OpenShift specific types to the json schemes
	appsv1.AddToScheme(scheme.Scheme)
	authorizationv1.AddToScheme(scheme.Scheme)
	buildv1.AddToScheme(scheme.Scheme)
	imagev1.AddToScheme(scheme.Scheme)
	networkv1.AddToScheme(scheme.Scheme)
	oauthv1.AddToScheme(scheme.Scheme)
	projectv1.AddToScheme(scheme.Scheme)
	quotav1.AddToScheme(scheme.Scheme)
	routev1.AddToScheme(scheme.Scheme)
	securityv1.AddToScheme(scheme.Scheme)
	templatev1.AddToScheme(scheme.Scheme)
	userv1.AddToScheme(scheme.Scheme)

	lib.CLog("info", "Discovering resources from OpenShift Namespace "+namespace)

	routes, _ := c.GetAllRoutes(namespace)
	for _, y := range routes.Items {
		y := y
		var rl ResourceList
		rl.Namespace = namespace
		//fmt.Println(y.TypeMeta.Kind)
		rl.Kind = "Route"
		rl.Name = y.ObjectMeta.Name
		rl.APIVersion = y.TypeMeta.APIVersion
		rl.UID = y.ObjectMeta.UID
		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)
		serializer := k8sjson.NewSerializerWithOptions(k8sjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme,
			k8sjson.SerializerOptions{
				Yaml:   yaml,
				Pretty: true,
				Strict: true,
			},
		)
		err = serializer.Encode(&y, writer)
		if err != nil {
			lib.CLog("error", "Building resource list for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			lib.CLog("error", "Unable to write to buffer for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}

		rl.Payload = *buff
		rl.PayloadStr = rl.Payload.String()
		resourcelist = append(resourcelist, rl)
	}

	services, _ := c.GetAllServices(namespace)
	for _, y := range services.Items {
		y := y
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = "Service"
		rl.Name = y.ObjectMeta.Name
		rl.APIVersion = y.TypeMeta.APIVersion
		rl.UID = y.ObjectMeta.UID
		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)
		serializer := k8sjson.NewSerializerWithOptions(k8sjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme,
			k8sjson.SerializerOptions{
				Yaml:   yaml,
				Pretty: true,
				Strict: true,
			},
		)
		err := serializer.Encode(&y, writer)
		if err != nil {
			lib.CLog("error", "Building resource list for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			lib.CLog("error", "Unable to write to buffer for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	deploymentconfigs, _ := c.GetAllDeploymentConfigs(namespace)
	for _, y := range deploymentconfigs.Items {
		y := y
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = "DeploymentConfig"
		rl.Name = y.ObjectMeta.Name
		rl.APIVersion = y.TypeMeta.APIVersion
		rl.UID = y.ObjectMeta.UID
		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)
		serializer := k8sjson.NewSerializerWithOptions(k8sjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme,
			k8sjson.SerializerOptions{
				Yaml:   yaml,
				Pretty: true,
				Strict: true,
			},
		)
		err := serializer.Encode(&y, writer)
		if err != nil {
			lib.CLog("error", "Building resource list for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			lib.CLog("error", "Unable to write to buffer for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}
		rl.Payload = *buff
		rl.PayloadStr = rl.Payload.String()
		resourcelist = append(resourcelist, rl)
	}

	deployments, _ := c.GetAllDeployments(namespace)
	for _, y := range deployments.Items {
		y := y
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = "Deployment"
		rl.Name = y.ObjectMeta.Name
		rl.APIVersion = y.TypeMeta.APIVersion
		rl.UID = y.ObjectMeta.UID
		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)
		serializer := k8sjson.NewSerializerWithOptions(k8sjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme,
			k8sjson.SerializerOptions{
				Yaml:   yaml,
				Pretty: true,
				Strict: true,
			},
		)
		err := serializer.Encode(&y, writer)
		if err != nil {
			lib.CLog("error", "Building resource list for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			lib.CLog("error", "Unable to write to buffer for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}
		rl.Payload = *buff
		rl.PayloadStr = rl.Payload.String()
		resourcelist = append(resourcelist, rl)
	}

	serviceaccounts, _ := c.GetAllServiceAccounts(namespace)
	for _, y := range serviceaccounts.Items {
		y := y
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = "ServiceAccount"
		rl.Name = y.ObjectMeta.Name
		rl.APIVersion = y.TypeMeta.APIVersion
		rl.UID = y.ObjectMeta.UID
		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)
		serializer := k8sjson.NewSerializerWithOptions(k8sjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme,
			k8sjson.SerializerOptions{
				Yaml:   yaml,
				Pretty: true,
				Strict: true,
			},
		)
		err := serializer.Encode(&y, writer)
		if err != nil {
			lib.CLog("error", "Building resource list for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			lib.CLog("error", "Unable to write to buffer for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	statefulsets, _ := c.GetAllStatefulSets(namespace)
	for _, y := range statefulsets.Items {
		y := y
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = "StatefulSet"
		rl.Name = y.ObjectMeta.Name
		rl.APIVersion = y.TypeMeta.APIVersion
		rl.UID = y.ObjectMeta.UID
		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)
		serializer := k8sjson.NewSerializerWithOptions(k8sjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme,
			k8sjson.SerializerOptions{
				Yaml:   yaml,
				Pretty: true,
				Strict: true,
			},
		)
		err := serializer.Encode(&y, writer)
		if err != nil {
			lib.CLog("error", "Building resource list for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			lib.CLog("error", "Unable to write to buffer for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	build, _ := c.GetAllBuilds(namespace)
	for _, y := range build.Items {
		y := y
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = "Build"
		rl.Name = y.ObjectMeta.Name
		rl.APIVersion = y.TypeMeta.APIVersion
		rl.UID = y.ObjectMeta.UID
		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)
		serializer := k8sjson.NewSerializerWithOptions(k8sjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme,
			k8sjson.SerializerOptions{
				Yaml:   yaml,
				Pretty: true,
				Strict: true,
			},
		)
		err := serializer.Encode(&y, writer)
		if err != nil {
			lib.CLog("error", "Building resource list for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			lib.CLog("error", "Unable to write to buffer for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	configmap, _ := c.GetAllConfigMaps(namespace)
	for _, y := range configmap.Items {
		y := y
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = "ConfigMap"
		rl.Name = y.ObjectMeta.Name
		rl.APIVersion = y.TypeMeta.APIVersion
		rl.UID = y.ObjectMeta.UID
		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)
		serializer := k8sjson.NewSerializerWithOptions(k8sjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme,
			k8sjson.SerializerOptions{
				Yaml:   yaml,
				Pretty: true,
				Strict: true,
			},
		)
		err := serializer.Encode(&y, writer)
		if err != nil {
			lib.CLog("error", "Building resource list for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			lib.CLog("error", "Unable to write to buffer for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	imagestream, _ := c.GetAllImageStreams(namespace)
	for _, y := range imagestream.Items {
		y := y
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = "ImageStream"
		rl.Name = y.ObjectMeta.Name
		rl.APIVersion = y.TypeMeta.APIVersion
		rl.UID = y.ObjectMeta.UID
		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)
		serializer := k8sjson.NewSerializerWithOptions(k8sjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme,
			k8sjson.SerializerOptions{
				Yaml:   yaml,
				Pretty: true,
				Strict: true,
			},
		)
		err := serializer.Encode(&y, writer)
		if err != nil {
			lib.CLog("error", "Building resource list for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			lib.CLog("error", "Unable to write to buffer for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	templates, _ := c.GetAllTemplates(namespace)
	for _, y := range templates.Items {
		y := y
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = "Template"
		rl.Name = y.ObjectMeta.Name
		rl.APIVersion = y.TypeMeta.APIVersion
		rl.UID = y.ObjectMeta.UID
		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)
		serializer := k8sjson.NewSerializerWithOptions(k8sjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme,
			k8sjson.SerializerOptions{
				Yaml:   yaml,
				Pretty: true,
				Strict: true,
			},
		)
		err := serializer.Encode(&y, writer)
		if err != nil {
			lib.CLog("error", "Building resource list for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			lib.CLog("error", "Unable to write to buffer for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	jobs, _ := c.GetAllJobs(namespace)
	for _, y := range jobs.Items {
		y := y
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = "Job"
		rl.Name = y.ObjectMeta.Name
		rl.APIVersion = y.TypeMeta.APIVersion
		rl.UID = y.ObjectMeta.UID
		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)
		serializer := k8sjson.NewSerializerWithOptions(k8sjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme,
			k8sjson.SerializerOptions{
				Yaml:   yaml,
				Pretty: true,
				Strict: true,
			},
		)
		err := serializer.Encode(&y, writer)
		if err != nil {
			lib.CLog("error", "Building resource list for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			lib.CLog("error", "Unable to write to buffer for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	pvc, _ := c.GetAllPVC(namespace)
	for _, y := range pvc.Items {
		y := y
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = "PersistentVolumeClaim"
		rl.Name = y.ObjectMeta.Name
		rl.APIVersion = y.TypeMeta.APIVersion
		rl.UID = y.ObjectMeta.UID
		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)
		serializer := k8sjson.NewSerializerWithOptions(k8sjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme,
			k8sjson.SerializerOptions{
				Yaml:   yaml,
				Pretty: true,
				Strict: true,
			},
		)
		err := serializer.Encode(&y, writer)
		if err != nil {
			lib.CLog("error", "Building resource list for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			lib.CLog("error", "Unable to write to buffer for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	pv, _ := c.GetAllPV(namespace)
	for _, y := range pv.Items {
		y := y
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = "PersistentVolume"
		rl.Name = y.ObjectMeta.Name
		rl.APIVersion = y.TypeMeta.APIVersion
		rl.UID = y.ObjectMeta.UID
		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)
		serializer := k8sjson.NewSerializerWithOptions(k8sjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme,
			k8sjson.SerializerOptions{
				Yaml:   yaml,
				Pretty: true,
				Strict: true,
			},
		)
		err := serializer.Encode(&y, writer)
		if err != nil {
			lib.CLog("error", "Building resource list for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			lib.CLog("error", "Unable to write to buffer for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	secret, _ := c.GetAllSecrets(namespace)
	for _, y := range secret.Items {
		y := y
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = "Secret"
		rl.Name = y.ObjectMeta.Name
		rl.APIVersion = y.TypeMeta.APIVersion
		rl.UID = y.ObjectMeta.UID
		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)
		serializer := k8sjson.NewSerializerWithOptions(k8sjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme,
			k8sjson.SerializerOptions{
				Yaml:   yaml,
				Pretty: true,
				Strict: true,
			},
		)
		err := serializer.Encode(&y, writer)
		if err != nil {
			lib.CLog("error", "Building resource list for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			lib.CLog("error", "Unable to write to buffer for object: "+rl.Kind+" "+rl.Name, err)
			rl.Error = err
		}
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}
	return resourcelist, err
}
