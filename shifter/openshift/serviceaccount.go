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

// Get all Service Accounts by Namespace
func (c Openshift) GetAllServiceAccounts(namespace string) (*v1.ServiceAccountList, error) {
	// Uses KNative Structs
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &v1.ServiceAccountList{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get All OpenShift Service Accounts By Namespace
	serviceAccounts, err := cluster.CoreV1().ServiceAccounts(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		// Error: Getting All OpenShift Service Accounts By Namespace
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Service Accounts from Namespace: '%s'. Error Text: '%s'. ", namespace, err.Error())
		return &v1.ServiceAccountList{}, err
	} else {
		// Success: Getting All OpenShift Service Accounts By Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting OpenShift Service Accounts from Namespace: '%s'.", namespace)
		// Return Service Accounts
		return serviceAccounts, err
	}

}

// Get OpenShift Service Account by Name from Namespace
func (c Openshift) GetServiceAccount(name string, namespace string) (*v1.ServiceAccount, error) {
	// Uses KNative Structs
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &v1.ServiceAccount{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get OpenShift Service Account By Name from Namespace
	serviceAccount, err := cluster.CoreV1().ServiceAccounts(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		// Error: Getting OpenShift Service Account By Name & Namespace
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Service Account with Name: '%s' from Namespace: '%s'. Error Text: '%s'. ", name, namespace, err.Error())
		return &v1.ServiceAccount{}, err
	} else {
		// Success: Getting OpenShift Service Account By Name & Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting OpenShift Service Account with Name: '%s' from Namespace: '%s'.", name, namespace)
		// Return Service Account
		return serviceAccount, err
	}
}
