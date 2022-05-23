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
	restclientcmdapi "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"log"
)

type Openshift struct {
	Endpoint  string
	AuthToken string
<<<<<<< HEAD
<<<<<<< HEAD
	Username  string
	Password  string
=======
>>>>>>> 6f965c1 (Moving to native go client)
=======
	Username string
	Password string
>>>>>>> def3249 (Fixing cluster config from the api)
}

func (cluster *Openshift) clusterClient() *restclientcmdapi.Config {
	config := clientcmdapi.NewConfig()
	config.Clusters["cluster"] = &clientcmdapi.Cluster{
		InsecureSkipTLSVerify: true,
		Server:                cluster.Endpoint,
	}

	config.AuthInfos["cluster-auth"] = &clientcmdapi.AuthInfo{
<<<<<<< HEAD
		Token:    cluster.AuthToken,
=======
		Token: cluster.AuthToken,
>>>>>>> def3249 (Fixing cluster config from the api)
		Username: cluster.Username,
		Password: cluster.Password,
	}

	config.Contexts["ctx"] = &clientcmdapi.Context{
		Cluster:  "cluster",
		AuthInfo: "cluster-auth",
	}

	config.CurrentContext = "ctx"
	clusterConfig := clientcmd.NewNonInteractiveClientConfig(*config, "ctx", &clientcmd.ConfigOverrides{}, nil)

	c, err := clusterConfig.ClientConfig()
	if err != nil {
		log.Println(err.Error())
	}

	return c
}
