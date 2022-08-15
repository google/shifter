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

	v1 "github.com/openshift/api/project/v1"
	projectv1 "github.com/openshift/client-go/project/clientset/versioned/typed/project/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Get all OpenShift Projects by Namespace
func (c Openshift) GetAllProjects() (*v1.ProjectList, error) {
	// Uses Custom OpenShift Structs
	cluster, err := projectv1.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &v1.ProjectList{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get All OpenShift Projects By Namespace
	projectList, err := cluster.Projects().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		// Error: Getting All OpenShift Projects By Namespace
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Projects. Error Text: '%s'. ", err.Error())
		return &v1.ProjectList{}, err
	} else {
		// Success: Getting All OpenShift Projects By Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting OpenShift Projects from Cluster.")
		// Return Projects
		return projectList, err
	}
}

// Get OpenShift Project by Name from Namespace
func (c Openshift) GetProject(name string) (*v1.Project, error) {
	// Uses Custom OpenShift Structs
	cluster, err := projectv1.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &v1.Project{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get OpenShift Project By Name from Namespace
	project, err := cluster.Projects().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		// Error: Getting OpenShift Project By Name & Namespace
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Project with Name: '%s'. Error Text: '%s'. ", name, err.Error())
		return &v1.Project{}, err
	} else {
		// Success: Getting OpenShift Project By Name & Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting OpenShift Project with Name: '%s'.", name)
		// Return Project
		return project, err
	}
}
