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

// Get all OpenShift Persistent Volumes by Namespace
func (c Openshift) GetAllPV(namespace string) (*v1.PersistentVolumeList, error) {
	// Uses KNative Structs
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &v1.PersistentVolumeList{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get All OpenShift Persistent Volumes By Namespace
	pvs, err := cluster.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		// Error: Getting All OpenShift Persistent Volumes By Namespace
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Persistent Volumes from Namespace: '%s'. Error Text: '%s'. ", namespace, err.Error())
		return &v1.PersistentVolumeList{}, err
	} else {
		// Success: Getting All OpenShift Persistent Volumes By Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting OpenShift Persistent Volumes from Namespace: '%s'.", namespace)
		// Return Persistent Volumes
		return pvs, err
	}

}

// Get OpenShift Persistent Volume by Name from Namespace
func (c Openshift) GetPV(name string, namespace string) (*v1.PersistentVolume, error) {
	// Uses KNative Structs
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &v1.PersistentVolume{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get OpenShift Persistent Volume By Name from Namespace
	pv, err := cluster.CoreV1().PersistentVolumes().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		// Error: Getting OpenShift Persistent Volume By Name & Namespace
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Persistent Volume with Name: '%s' from Namespace: '%s'. Error Text: '%s'. ", name, namespace, err.Error())
		return &v1.PersistentVolume{}, err
	} else {
		// Success: Getting OpenShift Persistent Volume By Name & Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting OpenShift Persistent Volume with Name: '%s' from Namespace: '%s'.", name, namespace)
		// Return Persistent Volume
		return pv, err
	}

}

// Get all OpenShift Persistent Volume Claims by Namespace
func (c Openshift) GetAllPVC(namespace string) (*v1.PersistentVolumeClaimList, error) {
	// Uses KNative Structs
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &v1.PersistentVolumeClaimList{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get All OpenShift Persistent Volume Claims By Namespace
	pvcList, err := cluster.CoreV1().PersistentVolumeClaims(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		// Error: Getting All OpenShift Persistent Volume Claims By Namespace
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Persistent Volume Claims from Namespace: '%s'. Error Text: '%s'. ", namespace, err.Error())
		return &v1.PersistentVolumeClaimList{}, err
	} else {
		// Success: Getting All OpenShift Persistent Volume Claims By Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting OpenShift Persistent Volume Claims from Namespace: '%s'.", namespace)
		// Return Persistent Volume Claims
		return pvcList, err
	}

}

// Get OpenShift Persistent Volume Claim by Name from Namespace
func (c Openshift) GetPVC(name string, namespace string) (*v1.PersistentVolumeClaim, error) {
	// Uses KNative Structs
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &v1.PersistentVolumeClaim{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get OpenShift Persistent Volume Claim By Name from Namespace
	pvc, err := cluster.CoreV1().PersistentVolumeClaims(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		// Error: Getting OpenShift Persistent Volume Claim By Name & Namespace
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Persistent Volume Claim with Name: '%s' from Namespace: '%s'. Error Text: '%s'. ", name, namespace, err.Error())
		return &v1.PersistentVolumeClaim{}, err
	} else {
		// Success: Getting OpenShift Persistent Volume Claim By Name & Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting OpenShift Persistent Volume Claim with Name: '%s' from Namespace: '%s'.", name, namespace)
		// Return Persistent Volume Claim
		return pvc, err
	}

}
