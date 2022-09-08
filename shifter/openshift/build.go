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
	"context"
	"shifter/lib"

	v1 "github.com/openshift/api/build/v1"
	os "github.com/openshift/client-go/build/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c Openshift) GetAllBuilds(namespace string) (*v1.BuildList, error) {
	cluster, err := os.NewForConfig(c.clusterClient())
	if err != nil {
		lib.CLog("error", "Unable to connect to cluster", err)
		return &v1.BuildList{}, err
	}

	buildList, err := cluster.BuildV1().Builds(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		lib.CLog("error", "Getting Builds from Namespace: "+namespace, err)
		return &v1.BuildList{}, err
	}
	lib.CLog("debug", "Getting Builds from Namespace: "+namespace)
	return buildList, err
}

func (c Openshift) GetBuild(name string, namespace string) (*v1.Build, error) {
	cluster, err := os.NewForConfig(c.clusterClient())
	if err != nil {
		lib.CLog("error", "Unable to connect to cluster", err)
		return &v1.Build{}, err
	}

	build, err := cluster.BuildV1().Builds(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		lib.CLog("error", "Getting Build with Name: "+name+" from Namespace: "+namespace, err)
		return &v1.Build{}, err
	}
	lib.CLog("info", "Getting Build with Name: "+name+" from Namespace: "+namespace)

	return build, nil
}
