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

	v1 "github.com/openshift/api/build/v1"
	os "github.com/openshift/client-go/build/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Get all OpenShift Builds by Namespace
func (c Openshift) GetAllBuilds(namespace string) (*v1.BuildList, error) {
	// Uses Custom OpenShift Structs
	cluster, err := os.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &v1.BuildList{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get All OpenShift Builds By Namespace
	buildList, err := cluster.BuildV1().Builds(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		// Error: Getting All OpenShift Builds By Namespace
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Builds from Namespace: '%s'. Error Text: '%s'. ", namespace, err.Error())
		return &v1.BuildList{}, err
	} else {
		// Success: Getting All OpenShift Builds By Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting OpenShift Builds from Namespace: '%s'.", namespace)
		// Return Builds
		return buildList, err
	}

}

// Get OpenShift Build by Name from a Namespace
func (c Openshift) GetBuild(name string, namespace string) (*v1.Build, error) {
	// Uses Custom OpenShift Structs
	cluster, err := os.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &v1.Build{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get OpenShift Builds By Name from Namespace
	build, err := cluster.BuildV1().Builds(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		// Error: Getting OpenShift Build By Name & Namespace
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Build with Name: '%s' from Namespace: '%s'. Error Text: '%s'. ", name, namespace, err.Error())
		return &v1.Build{}, err
	} else {
		// Success: Getting OpenShift Build By Name & Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting OpenShift Build with Name: '%s' from Namespace: '%s'.", name, namespace)
		// Return Build
		return build, err
	}
}
