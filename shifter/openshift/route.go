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

	v1 "github.com/openshift/api/route/v1"
	os "github.com/openshift/client-go/route/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Get all OpenShift Routes by Namespace
func (c Openshift) GetAllRoutes(namespace string) (*v1.RouteList, error) {
	// Uses Custom OpenShift Structs
	cluster, err := os.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &v1.RouteList{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get All OpenShift Routes By Namespace
	routeList, err := cluster.RouteV1().Routes(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		// Error: Getting All OpenShift Routes By Namespace
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Routes from Namespace: '%s'. Error Text: '%s'. ", namespace, err.Error())
		return &v1.RouteList{}, err
	} else {
		// Success: Getting All OpenShift Routes By Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting OpenShift Routes from Namespace: '%s'.", namespace)
		// Return Routes
		return routeList, err
	}

}

// Get OpenShift Route by Name from Namespace
func (c Openshift) GetRoute(name string, namespace string) (*v1.Route, error) {
	// Uses Custom OpenShift Structs
	cluster, err := os.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &v1.Route{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get OpenShift Route By Name from Namespace
	route, err := cluster.RouteV1().Routes(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		// Error: Getting OpenShift Route By Name & Namespace
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Route with Name: '%s' from Namespace: '%s'. Error Text: '%s'. ", name, namespace, err.Error())
		return &v1.Route{}, err
	} else {
		// Success: Getting OpenShift Route By Name & Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting OpenShift Route with Name: '%s' from Namespace: '%s'.", name, namespace)
		// Return Route
		return route, err
	}
}
