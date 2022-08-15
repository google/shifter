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
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Get all OpenShift ConfigMaps by Namespace
func (c Openshift) GetAllConfigMaps(namespace string) (*v1.ConfigMapList, error) {
	// Uses KNative Structs
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &v1.ConfigMapList{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get All OpenShift ConfigMaps By Namespace
	configmapList, err := cluster.CoreV1().ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		// Error: Getting All OpenShift ConfigMaps By Namespace
		log.Printf("🧰 ❌ ERROR: Getting OpenShift ConfigMaps from Namespace: '%s'. Error Text: '%s'. ", namespace, err.Error())
		return &v1.ConfigMapList{}, err
	} else {
		// Success: Getting All OpenShift ConfigMaps By Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting OpenShift ConfigMaps from Namespace: '%s'.", namespace)
		// Return ConfigMaps
		return configmapList, err
	}

}

// Get OpenShift ConfigMap by Name from Namespace
func (c Openshift) GetConfigMap(name string, namespace string) (*v1.ConfigMap, error) {
	// Uses KNative Structs
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &v1.ConfigMap{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get OpenShift ConfigMap By Name from Namespace
	configmap, err := cluster.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		// Error: Getting OpenShift ConfigMap By Name & Namespace
		log.Printf("🧰 ❌ ERROR: Getting OpenShift ConfigMap with Name: '%s' from Namespace: '%s'. Error Text: '%s'. ", name, namespace, err.Error())
		return &v1.ConfigMap{}, err
	} else {
		// Success: Getting OpenShift ConfigMap By Name & Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting OpenShift ConfigMap with Name: '%s' from Namespace: '%s'.", name, namespace)
		// Return ConfigMap
		return configmap, err
	}
}
