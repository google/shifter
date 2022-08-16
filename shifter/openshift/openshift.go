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
	"log"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	k8sjson "k8s.io/apimachinery/pkg/runtime/serializer/json"
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

	//"shifter/lib"
	//"reflect"
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
	Error     error
}

const (
	SECRET                = "Secret"
	PERSISTENTVOLUME      = "PersistentVolume"
	PERSISTENTVOLUMECLAIM = "PersistentVolumeClaim"
	JOB                   = "Job"
	TEMPLATE              = "Template"
	IMAGESTREAM           = "ImageStream"
	CONFIGMAP             = "ConfigMap"
	BUILD                 = "Build"
	STATEFULSET           = "StatefulSet"
	SERVICEACCOUNT        = "ServiceAccount"
	DEPLOYMENT            = "Deployment"
	DEPLOYMENTCONFIG      = "DeploymentConfig"
	SERVICE               = "Service"
	ROUTE                 = "Route"
)

func (c *Openshift) clusterClient() *restclientcmdapi.Config {

	// Instantiating Kubernetes CLI Configuration
	log.Printf("🧰 💡 INFO: Instantiating Kubernetes CLI Config")
	config := clientcmdapi.NewConfig()
	config.Clusters["cluster"] = &clientcmdapi.Cluster{
		InsecureSkipTLSVerify: true,
		Server:                c.Endpoint,
	}

	// Validate the OpenShift Cluster Token
	log.Printf("🧰 💡 INFO: Validating the OpenShift Cluster Token")
	if len(c.AuthToken) <= 5 {
		// Error the OpenShift Cluster Token
		log.Printf("🧰 ❌ ERROR: Invalid OpenShift Token, Unable to continue. ")
		os.Exit(1) // TODO - RETURN ERROR
	} else {
		// Valid OpenShift Cluster Token Provided
		log.Printf("🧰 ✅ SUCCESS: Valid OpenShift Token Provided. ")
	}

	// Setting Kubernetes Config Cluster Auth
	log.Printf("🧰 💡 INFO: Setting Kubernetes Config Cluster Auth")
	config.AuthInfos["cluster-auth"] = &clientcmdapi.AuthInfo{
		Token:    c.AuthToken,
		Username: c.Username,
		Password: c.Password,
	}

	// Setting Kubernetes Config Contexts
	log.Printf("🧰 💡 INFO: Setting Kubernetes Config Contexts")
	config.Contexts["ctx"] = &clientcmdapi.Context{
		Cluster:  "cluster",
		AuthInfo: "cluster-auth",
	}

	// Setting Kubernetes Current Context = CTX
	config.CurrentContext = "ctx"
	log.Printf("🧰 💡 INFO: Creating New Non Interatcive Client Config for Kubernetes")
	clusterConfig := clientcmd.NewNonInteractiveClientConfig(*config, "ctx", &clientcmd.ConfigOverrides{}, nil)

	// Configuring the OpenShift Cluster Client
	log.Printf("🧰 💡 INFO: Configuring OpenShift Cluster Client")
	cl, err := clusterConfig.ClientConfig()
	if err != nil {
		// Error: Configuring OpenShift Cluster Client
		log.Printf("🧰 ❌ ERROR: Configuring OpenShift Cluster Client: '%s'. ", err.Error())
		os.Exit(1) // TODO - RETURN ERROR
	} else {
		// Success: Configuring OpenShift Cluster Client
		log.Printf("🧰 ✅ SUCCESS: OpenShift Cluster Client Configured.")
	}
	// Return Cluster Configuration
	return cl
}

func (c *Openshift) ExportNSResources(namespace string, outputPath string) error {
	var resourcelist []ResourceList

	if namespace != "" {
		resourcelist = c.getResources(namespace, true)
	} else {
		projects, err := c.GetAllProjects()
		if err != nil {
			// Error: Unable to get All OpenShift Project
			log.Printf("🧰 ❌ ERROR: Unable to get All OpenShift Project.")
			return err
		}
		for _, p := range projects.Items {
			rl := c.getResources(p.ObjectMeta.Name, true)
			for _, v := range rl {
				resourcelist = append(resourcelist, v)
			}
		}
	}
	for _, res := range resourcelist {
		// TODO - Hard Coded Values (YAML - LOCAL)
		log.Printf("🧰 💡 INFO: Exporting object: " + res.Namespace + "\\" + res.Name + " of kind " + res.Kind)
		fileObj := &ops.FileObject{
			StorageType:   "local",
			Path:          (outputPath + "/" + res.Namespace + "/" + res.Name),
			Ext:           "yaml",
			Content:       res.Payload,
			ContentLength: res.Payload.Len(),
		}
		err := fileObj.WriteFile()
		if err != nil {
			// Error: Error Writing File
			return err
		}
	}
	// All Successful.
	return nil
}

func (c *Openshift) ConvertNSResources(namespace string, flags map[string]string, outputPath string) error {
	var resourcelist []ResourceList

	if namespace != "" {
		resourcelist = c.getResources(namespace, false)
	} else {
		projects, err := c.GetAllProjects()
		if err != nil {
			// Error: Unable to get All OpenShift Project
			log.Printf("🧰 ❌ ERROR: Unable to get All OpenShift Project.")
			return err
		}
		for _, p := range projects.Items {
			rl := c.getResources(p.ObjectMeta.Name, false)
			for _, v := range rl {
				resourcelist = append(resourcelist, v)
			}
		}
	}

	for _, res := range resourcelist {
		log.Println("🧰 💡 INFO: Converting object: " + res.Namespace + "\\" + res.Name + " of kind " + res.Kind)
		// Create Shifter Processor
		obj, err := processor.Processor(res.Payload.Bytes(), res.Kind, flags)
		if err != nil {
			// Error: Unable to Create Shifter Processor
			log.Printf("🧰 ❌ ERROR: Create Shifter Processor.")
			return err
		} else {
			// Succes: Creating Shifter Processor
			log.Printf("🧰 ✅ SUCCESS: Shifter Processor Successufly Created.")
		}
		// Create Shifer Generator
		convertedObject, err := generator.NewGenerator("yaml", res.Name, obj)
		if err != nil {
			// Error: Unable to Create Shifter Generator
			log.Printf("🧰 ❌ ERROR: Create Shifter Generator.")
			return err
		} else {
			// Succes: Creating Shifter Generator
			log.Printf("🧰 ✅ SUCCESS: Shifter Generator Successufly Created.")
		}
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
				// Error: Error Writing File
				return err
			}
		}
	}
	return nil
}

func (c *Openshift) ListNSResources(csvoutput bool, namespace string) error {
	var resourcelist []ResourceList

	if namespace != "" {
		resourcelist = c.getResources(namespace, false)
	} else {
		projects, err := c.GetAllProjects()
		if err != nil {
			// Error: Unable to get All OpenShift Project
			log.Printf("🧰 ❌ ERROR: Unable to get All OpenShift Project.")
			return err
		}
		for _, p := range projects.Items {
			rl := c.getResources(p.ObjectMeta.Name, false)
			for _, v := range rl {
				resourcelist = append(resourcelist, v)
			}
		}
	}

	t := table.NewWriter()
	//t.SetStyle(table.StyleColoredBright)
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

	// All Successful
	return nil
}

func (c *Openshift) getResources(namespace string, yaml bool) []ResourceList {
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

	log.Println("🧰 💡 INFO: Discovering resources from OpenShift Namespace/Project " + namespace)

	routes, _ := c.GetAllRoutes(namespace)
	//log.Println(reflect.TypeOf(routes))
	for _, y := range routes.Items {
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = ROUTE
		rl.Name = y.ObjectMeta.Name
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
			// Error: Building Resource List
			log.Printf("🧰 ❌ ERROR: Building Resource List: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			// Error: Writing to the Byte Buffer
			log.Printf("🧰 ❌ ERROR: Writing data to memory: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}

		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	services, _ := c.GetAllServices(namespace)
	for _, y := range services.Items {
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = SERVICE
		rl.Name = y.ObjectMeta.Name
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
			// Error: Building Resource List
			log.Printf("🧰 ❌ ERROR: Building Resource List: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			// Error: Writing to the Byte Buffer
			log.Printf("🧰 ❌ ERROR: Writing data to memory: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	deploymentconfigs, _ := c.GetAllDeploymentConfigs(namespace)
	for _, y := range deploymentconfigs.Items {
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = DEPLOYMENTCONFIG
		rl.Name = y.ObjectMeta.Name
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
			// Error: Building Resource List
			log.Printf("🧰 ❌ ERROR: Building Resource List: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			// Error: Writing to the Byte Buffer
			log.Printf("🧰 ❌ ERROR: Writing data to memory: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	deployments, _ := c.GetAllDeployments(namespace)
	for _, y := range deployments.Items {
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = DEPLOYMENT
		rl.Name = y.ObjectMeta.Name
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
			// Error: Building Resource List
			log.Printf("🧰 ❌ ERROR: Building Resource List: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			// Error: Writing to the Byte Buffer
			log.Printf("🧰 ❌ ERROR: Writing data to memory: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	serviceaccounts, _ := c.GetAllServiceAccounts(namespace)
	for _, y := range serviceaccounts.Items {
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = SERVICEACCOUNT
		rl.Name = y.ObjectMeta.Name
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
			// Error: Building Resource List
			log.Printf("🧰 ❌ ERROR: Building Resource List: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			// Error: Writing to the Byte Buffer
			log.Printf("🧰 ❌ ERROR: Writing data to memory: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	statefulsets, _ := c.GetAllStatefulSets(namespace)
	for _, y := range statefulsets.Items {
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = STATEFULSET
		rl.Name = y.ObjectMeta.Name
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
			// Error: Building Resource List
			log.Printf("🧰 ❌ ERROR: Building Resource List: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			// Error: Writing to the Byte Buffer
			log.Printf("🧰 ❌ ERROR: Writing data to memory: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	build, _ := c.GetAllBuilds(namespace)
	for _, y := range build.Items {
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = BUILD
		rl.Name = y.ObjectMeta.Name
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
			// Error: Building Resource List
			log.Printf("🧰 ❌ ERROR: Building Resource List: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			// Error: Writing to the Byte Buffer
			log.Printf("🧰 ❌ ERROR: Writing data to memory: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	configmap, _ := c.GetAllConfigMaps(namespace)
	for _, y := range configmap.Items {
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = CONFIGMAP
		rl.Name = y.ObjectMeta.Name
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
			// Error: Building Resource List
			log.Printf("🧰 ❌ ERROR: Building Resource List: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			// Error: Writing to the Byte Buffer
			log.Printf("🧰 ❌ ERROR: Writing data to memory: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	imagestream, _ := c.GetAllImageStreams(namespace)
	for _, y := range imagestream.Items {
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = IMAGESTREAM
		rl.Name = y.ObjectMeta.Name
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
			// Error: Building Resource List
			log.Printf("🧰 ❌ ERROR: Building Resource List: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			// Error: Writing to the Byte Buffer
			log.Printf("🧰 ❌ ERROR: Writing data to memory: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	templates, _ := c.GetAllTemplates(namespace)
	for _, y := range templates.Items {
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = TEMPLATE
		rl.Name = y.ObjectMeta.Name
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
			// Error: Building Resource List
			log.Printf("🧰 ❌ ERROR: Building Resource List: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			// Error: Writing to the Byte Buffer
			log.Printf("🧰 ❌ ERROR: Writing data to memory: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	jobs, _ := c.GetAllJobs(namespace)
	for _, y := range jobs.Items {
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = JOB
		rl.Name = y.ObjectMeta.Name
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
			// Error: Building Resource List
			log.Printf("🧰 ❌ ERROR: Building Resource List: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			// Error: Writing to the Byte Buffer
			log.Printf("🧰 ❌ ERROR: Writing data to memory: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	pvc, _ := c.GetAllPVC(namespace)
	for _, y := range pvc.Items {
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = PERSISTENTVOLUMECLAIM
		rl.Name = y.ObjectMeta.Name
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
			// Error: Building Resource List
			log.Printf("🧰 ❌ ERROR: Building Resource List: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			// Error: Writing to the Byte Buffer
			log.Printf("🧰 ❌ ERROR: Writing data to memory: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	pv, _ := c.GetAllPV(namespace)
	for _, y := range pv.Items {
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = PERSISTENTVOLUME
		rl.Name = y.ObjectMeta.Name
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
			// Error: Building Resource List
			log.Printf("🧰 ❌ ERROR: Building Resource List: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			// Error: Writing to the Byte Buffer
			log.Printf("🧰 ❌ ERROR: Writing data to memory: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}

	secret, _ := c.GetAllSecrets(namespace)
	for _, y := range secret.Items {
		var rl ResourceList
		rl.Namespace = namespace
		rl.Kind = SECRET
		rl.Name = y.ObjectMeta.Name
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
			// Error: Building Resource List
			log.Printf("🧰 ❌ ERROR: Building Resource List: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}
		err = writer.Flush()
		if err != nil {
			// Error: Writing to the Byte Buffer
			log.Printf("🧰 ❌ ERROR: Writing data to memory: '%s'. ", err.Error())
			log.Printf("🧰 ❌ ERROR: Error Object Kind: '%s' and Name: '%s'. ", rl.Kind, rl.Name)
			rl.Error = err
		}
		rl.Payload = *buff
		resourcelist = append(resourcelist, rl)
	}
	return resourcelist
}
