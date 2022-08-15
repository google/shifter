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

	osappsv1 "github.com/openshift/api/apps/v1"
	os "github.com/openshift/client-go/apps/clientset/versioned"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Get all OpenShift Deployments by Namespace
func (c Openshift) GetAllDeployments(namespace string) (*v1.DeploymentList, error) {
	// Uses KNative Structs
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &v1.DeploymentList{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get All OpenShift Deployments By Namespace
	deploymentList, err := cluster.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		// Error: Getting All OpenShift Deployments By Namespace
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Deployments from Namespace: '%s'. Error Text: '%s'. ", namespace, err.Error())
		return &v1.DeploymentList{}, err
	} else {
		// Success: Getting All OpenShift Deployments By Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting OpenShift Deployments from Namespace: '%s'.", namespace)
		// Return Deployments
		return deploymentList, err
	}
}

// Get OpenShift Deployment by Name from Namespace
func (c Openshift) GetDeployment(name string, namespace string) (*v1.Deployment, error) {
	// Uses KNative Structs
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &v1.Deployment{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get OpenShift Deployment By Name from Namespace
	deployment, err := cluster.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		// Error: Getting OpenShift Deployment By Name & Namespace
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Deployment with Name: '%s' from Namespace: '%s'. Error Text: '%s'. ", name, namespace, err.Error())
		return &v1.Deployment{}, err
	} else {
		// Success: Getting OpenShift Deployment By Name & Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting OpenShift Deployment with Name: '%s' from Namespace: '%s'.", name, namespace)
		// Return Deployment
		return deployment, err
	}
}

// Get All OpenShift DeploymentConfigs by Namespace
func (c Openshift) GetAllDeploymentConfigs(namespace string) (*osappsv1.DeploymentConfigList, error) {
	// Uses Custom OpenShift Structs
	cluster, err := os.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &osappsv1.DeploymentConfigList{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get All Deployment Configs from Namespace
	depCfgLst, err := cluster.AppsV1().DeploymentConfigs(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		// Error: Getting All OpenShift Deployment Configurations By Namespace
		log.Printf("🧰 ❌ ERROR: Getting All OpenShift Deployment Configurations by Namespace: '%s'. Error Text: '%s'. ", namespace, err.Error())
		return &osappsv1.DeploymentConfigList{}, err
	} else {
		// Success: Getting All OpenShift Deployment Configurations By Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting All OpenShift Deployment Configurations by Namespace: '%s'.", namespace)
		// Return Deployment Configs
		return depCfgLst, err
	}
}

// Get OpenShift DeploymentConfig by Name from Namespace
func (c Openshift) GetDeploymentConfig(namespace string, name string) (*osappsv1.DeploymentConfig, error) {
	// Uses Custom OpenShift Structs
	cluster, err := os.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &osappsv1.DeploymentConfig{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get All Deployment Config By Name from Namespace
	depCfg, err := cluster.AppsV1().DeploymentConfigs(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		// Error: Getting OpenShift Deployment Configuration By Name & Namespace
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Deployment Configuration with Name: '%s' from Namespace: '%s'. Error Text: '%s'. ", name, namespace, err.Error())
		return &osappsv1.DeploymentConfig{}, err
	} else {
		// Success: Getting OpenShift Deployment Configuration By Name & Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting OpenShift Deployment Configuration with Name: '%s' from Namespace: '%s'.", name, namespace)
		// Return Deployment Config
		return depCfg, err
	}
}
