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
	"fmt"
	restclientcmdapi "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"log"
)

type Openshift struct {
	Endpoint  string
	AuthToken string
	Username  string
	Password  string
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

func (c *Openshift) ListAllResources(namespace string) {
	projects, err := c.GetAllProjects()
	if err != nil {
		fmt.Println(err)
	}
	for _, y := range projects.Items {
		fmt.Println(y.ObjectMeta.Name)
		fmt.Println("|_")

		routes, _ := c.GetAllRoutes(y.ObjectMeta.Name)
		for _, r := range routes.Items {
			fmt.Println("  [Route] " + r.ObjectMeta.Name)
		}

		services, _ := c.GetAllServices(y.ObjectMeta.Name)
		for _, s := range services.Items {
			fmt.Println("  [Service] " + s.ObjectMeta.Name)
		}

		deploymentconfigs, _ := c.GetAllDeploymentConfigs(y.ObjectMeta.Name)
		for _, d := range deploymentconfigs.Items {
			fmt.Println("  [DeploymentConfig] " + d.ObjectMeta.Name)
		}

		build, _ := c.GetAllBuilds(y.ObjectMeta.Name)
		for _, b := range build.Items {
			fmt.Println("  [Build] " + b.ObjectMeta.Name)
		}

		configmap, _ := c.GetAllConfigMaps(y.ObjectMeta.Name)
		for _, c := range configmap.Items {
			fmt.Println("  [ConfigMap] " + c.ObjectMeta.Name)
		}

		image, _ := c.GetAllImages(y.ObjectMeta.Name)
		for _, i := range image.Items {
			fmt.Println("  [Image] " + i.ObjectMeta.Name)
		}

	}
}
