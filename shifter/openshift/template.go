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

	v1 "github.com/openshift/api/template/v1"
	templatev1 "github.com/openshift/client-go/template/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Get all OpenShift Template by Namespace
func (c Openshift) GetAllTemplates(namespace string) (*v1.TemplateList, error) {
	// Uses Custom OpenShift Structs
	cluster, err := templatev1.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &v1.TemplateList{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get All OpenShift Templates By Namespace
	templates, err := cluster.TemplateV1().Templates(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		// Error: Getting All OpenShift Templates By Namespace
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Templates from Namespace: '%s'. Error Text: '%s'. ", namespace, err.Error())
		return &v1.TemplateList{}, err
	} else {
		// Success: Getting All OpenShift Templates By Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting OpenShift Templates from Namespace: '%s'.", namespace)
		// Return Templates
		return templates, err
	}

}

// Get OpenShift Template by Name from Namespace
func (c Openshift) GetTemplate(name string, namespace string) (*v1.Template, error) {
	// Uses Custom OpenShift Structs
	cluster, err := templatev1.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &v1.Template{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get OpenShift Template By Name from Namespace
	template, err := cluster.TemplateV1().Templates(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		// Error: Getting OpenShift Template By Name & Namespace
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Template with Name: '%s' from Namespace: '%s'. Error Text: '%s'. ", name, namespace, err.Error())
		return &v1.Template{}, err
	} else {
		// Success: Getting OpenShift Template By Name & Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting OpenShift Template with Name: '%s' from Namespace: '%s'.", name, namespace)
		// Return Template
		return template, err
	}

}
