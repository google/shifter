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

	v1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Get all OpenShift Jobs by Namespace
func (c Openshift) GetAllJobs(namespace string) (*v1.JobList, error) {
	// Uses KNative Structs
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &v1.JobList{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get All OpenShift Jobs By Namespace
	jobList, err := cluster.BatchV1().Jobs(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		// Error: Getting All OpenShift Jobs By Namespace
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Jobs from Namespace: '%s'. Error Text: '%s'. ", namespace, err.Error())
		return &v1.JobList{}, err
	} else {
		// Success: Getting All OpenShift Jobs By Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting OpenShift Jobs from Namespace: '%s'.", namespace)
		// Return Jobs
		return jobList, err
	}

}

// Get OpenShift Job by Name from Namespace
func (c Openshift) GetJob(name string, namespace string) (*v1.Job, error) {
	// Uses KNative Structs
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &v1.Job{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get OpenShift Job By Name from Namespace
	job, err := cluster.BatchV1().Jobs(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		// Error: Getting OpenShift Job By Name & Namespace
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Job with Name: '%s' from Namespace: '%s'. Error Text: '%s'. ", name, namespace, err.Error())
		return &v1.Job{}, err
	} else {
		// Success: Getting OpenShift Job By Name & Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting OpenShift Job with Name: '%s' from Namespace: '%s'.", name, namespace)
		// Return Job
		return job, err
	}

}
