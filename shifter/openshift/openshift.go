/*
copyright 2019 google llc
licensed under the apache license, version 2.0 (the "license");
you may not use this file except in compliance with the license.
you may obtain a copy of the license at
    http://www.apache.org/licenses/license-2.0
unless required by applicable law or agreed to in writing, software
distributed under the license is distributed on an "as is" basis,
without warranties or conditions of any kind, either express or implied.
see the license for the specific language governing permissions and
limitations under the license.
*/

package openshift

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	k8sjson "k8s.io/apimachinery/pkg/runtime/serializer/json"
	restclientcmdapi "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"log"
	"os"

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

	"k8s.io/client-go/kubernetes/scheme"
	"shifter/processor"
	"shifter/ops"
)

type Openshift struct {
	Endpoint  string
	AuthToken string
	Username  string
	Password  string
}

type ResourceList struct {
	Namespace string
	Kind      string
	Name      string
	Payload   bytes.Buffer
}

func (c *Openshift) clusterClient() *restclientcmdapi.Config {
	config := clientcmdapi.NewConfig()
	config.Clusters["cluster"] = &clientcmdapi.Cluster{
		InsecureSkipTLSVerify: true,
		Server:                c.Endpoint,
	}

	config.AuthInfos["cluster-auth"] = &clientcmdapi.AuthInfo{
		Token:    c.AuthToken,
		Username: c.Username,
		Password: c.Password,
	}

	config.Contexts["ctx"] = &clientcmdapi.Context{
		Cluster:  "cluster",
		AuthInfo: "cluster-auth",
	}

	config.CurrentContext = "ctx"
	clusterConfig := clientcmd.NewNonInteractiveClientConfig(*config, "ctx", &clientcmd.ConfigOverrides{}, nil)

	cl, err := clusterConfig.ClientConfig()
	if err != nil {
		log.Println(err.Error())
	}

	return cl
}

func (c *Openshift) ExportNSResources(namespace string) {
	var resourcelist []ResourceList

	if namespace != "" {
		resourcelist = c.getResources(namespace)
	} else {
		projects, err := c.GetAllProjects()
		if err != nil {
			log.Println(err)
		}
		for _, p := range projects.Items {
			rl := c.getResources(p.ObjectMeta.Name)
			for _, v := range rl {
				resourcelist = append(resourcelist, v)
			}
		}
	}
	for _, obj := range resourcelist {
		//obj := processor.Processor(obj.Payload.Bytes(), obj.Kind, nil)
		log.Println(obj.Kind, obj.Name)

	}
}

func (c *Openshift) ConvertNSResources(namespace string, flags map[string]string) {
	var resourcelist []ResourceList

	if namespace != "" {
		resourcelist = c.getResources(namespace)
	} else {
		projects, err := c.GetAllProjects()
		if err != nil {
			log.Println(err)
		}
		for _, p := range projects.Items {
			rl := c.getResources(p.ObjectMeta.Name)
			for _, v := range rl {
				resourcelist = append(resourcelist, v)
			}
		}
	}

	for _, obj := range resourcelist {
		//obj := processor.Processor(obj.Payload.Bytes(), obj.Kind, nil)
		fmt.Println(obj.Kind)
		fmt.Println(obj.Name)
		fmt.Println(obj.Payload.String())
		test := obj.Payload
		obj := processor.Processor(test.Bytes(), obj.Kind, flags)
		fmt.Println(obj)
	}

}

func (c *Openshift) ListNSResources(csvoutput bool, namespace string) {
	var resourcelist []ResourceList

	if namespace != "" {
		resourcelist = c.getResources(namespace)
	} else {
		projects, err := c.GetAllProjects()
		if err != nil {
			log.Println(err)
		}
		for _, p := range projects.Items {
			rl := c.getResources(p.ObjectMeta.Name)
			for _, v := range rl {
				resourcelist = append(resourcelist, v)
			}
		}
	}

	t := table.NewWriter()
	//t.SetStyle(table.StyleColoredBright)
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Project", "Kind", "Name", "Payload"})

	for _, x := range resourcelist {
		t.AppendRow(table.Row{x.Namespace, x.Kind, x.Name, len(x.Payload.String())})
	}
	if csvoutput == false {
		t.Render()
	} else {
		t.RenderCSV()
	}
}

func (c *Openshift) getResources(namespace string) []ResourceList {
	var resourcelist []ResourceList

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

	routes, _ := c.GetAllRoutes(namespace)
	for _, y := range routes.Items {
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = "Route"
		rl.Name = y.ObjectMeta.Name
		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)
		yaml := k8sjson.NewYAMLSerializer(k8sjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme)
		err := yaml.Encode(&y, writer)
		if err != nil {
			log.Println(err)
		}
		writer.Flush()
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	services, _ := c.GetAllServices(namespace)
	for _, y := range services.Items {
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = "Service"
		rl.Name = y.ObjectMeta.Name
		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)
		yaml := k8sjson.NewYAMLSerializer(k8sjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme)
		err := yaml.Encode(&y, writer)
		if err != nil {
			log.Println(err)
		}
		writer.Flush()
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	deploymentconfigs, _ := c.GetAllDeploymentConfigs(namespace)
	for _, y := range deploymentconfigs.Items {
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = "DeploymentConfig"
		rl.Name = y.ObjectMeta.Name
		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)
		yaml := k8sjson.NewYAMLSerializer(k8sjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme)
		err := yaml.Encode(&y, writer)
		if err != nil {
			log.Println(err)
		}
		writer.Flush()
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	deployments, _ := c.GetAllDeployments(namespace)
	for _, y := range deployments.Items {
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = "Deployment"
		rl.Name = y.ObjectMeta.Name
		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)
		yaml := k8sjson.NewYAMLSerializer(k8sjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme)
		err := yaml.Encode(&y, writer)
		if err != nil {
			log.Println(err)
		}
		writer.Flush()
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	serviceaccounts, _ := c.GetAllServiceAccounts(namespace)
	for _, y := range serviceaccounts.Items {
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = "ServiceAccount"
		rl.Name = y.ObjectMeta.Name
		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)
		yaml := k8sjson.NewYAMLSerializer(k8sjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme)
		err := yaml.Encode(&y, writer)
		if err != nil {
			log.Println(err)
		}
		writer.Flush()
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	statefulsets, _ := c.GetAllStatefulSets(namespace)
	for _, y := range statefulsets.Items {
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = "StatefulSet"
		rl.Name = y.ObjectMeta.Name
		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)
		yaml := k8sjson.NewYAMLSerializer(k8sjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme)
		err := yaml.Encode(&y, writer)
		if err != nil {
			log.Println(err)
		}
		writer.Flush()
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	build, _ := c.GetAllBuilds(namespace)
	for _, y := range build.Items {
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = "Build"
		rl.Name = y.ObjectMeta.Name
		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)
		yaml := k8sjson.NewYAMLSerializer(k8sjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme)
		err := yaml.Encode(&y, writer)
		if err != nil {
			log.Println(err)
		}
		writer.Flush()
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	configmap, _ := c.GetAllConfigMaps(namespace)
	for _, y := range configmap.Items {
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = "ConfigMap"
		rl.Name = y.ObjectMeta.Name
		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)
		yaml := k8sjson.NewYAMLSerializer(k8sjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme)
		err := yaml.Encode(&y, writer)
		if err != nil {
			log.Println(err)
		}
		writer.Flush()
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	image, _ := c.GetAllImages(namespace)
	for _, y := range image.Items {
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = "Image"
		rl.Name = y.ObjectMeta.Name
		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)
		yaml := k8sjson.NewYAMLSerializer(k8sjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme)
		err := yaml.Encode(&y, writer)
		if err != nil {
			log.Println(err)
		}
		writer.Flush()
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	templates, _ := c.GetAllTemplates(namespace)
	for _, y := range templates.Items {
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = "Template"
		rl.Name = y.ObjectMeta.Name
		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)
		yaml := k8sjson.NewYAMLSerializer(k8sjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme)
		err := yaml.Encode(&y, writer)
		if err != nil {
			log.Println(err)
		}
		writer.Flush()
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	jobs, _ := c.GetAllJobs(namespace)
	for _, y := range jobs.Items {
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = "Job"
		rl.Name = y.ObjectMeta.Name
		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)
		yaml := k8sjson.NewYAMLSerializer(k8sjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme)
		err := yaml.Encode(&y, writer)
		if err != nil {
			log.Println(err)
		}
		writer.Flush()
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	return resourcelist
}
