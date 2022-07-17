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

// ConfigMaps are part of the core kubernetes api so we switch to using the upstream kubernetes client
import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
)

func (c Openshift) GetAllConfigMaps(namespace string) (*v1.ConfigMapList, error) {
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		log.Println(err)
		return &v1.ConfigMapList{}, err
	}

	configmapList, err := cluster.CoreV1().ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Println(err)
		return &v1.ConfigMapList{}, err
	}

	return configmapList, nil

}

func (c Openshift) GetConfigMap(name string, namespace string) (*v1.ConfigMap, error) {
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		log.Println(err)
		return &v1.ConfigMap{}, err
	}

	configmap, err := cluster.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		log.Println(err)
		return &v1.ConfigMap{}, err
	}

	return configmap, nil
}
