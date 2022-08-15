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

	v1 "github.com/openshift/api/image/v1"
	imagev1 "github.com/openshift/client-go/image/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Get all OpenShift Images by Namespace
func (c Openshift) GetAllImages(namespace string) (*v1.ImageList, error) {
	// Uses Custom OpenShift Structs
	cluster, err := imagev1.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &v1.ImageList{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get All OpenShift Images By Namespace
	images, err := cluster.ImageV1().Images().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		// Error: Getting All OpenShift Images By Namespace
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Images from Namespace: '%s'. Error Text: '%s'. ", namespace, err.Error())
		return &v1.ImageList{}, err
	} else {
		// Success: Getting All OpenShift Images By Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting OpenShift Images from Namespace: '%s'.", namespace)
		// Return Images
		return images, err
	}
}

// Get OpenShift Image by Name from Namespace
func (c Openshift) GetImage(name string, namespace string) (*v1.Image, error) {
	// Uses Custom OpenShift Structs
	cluster, err := imagev1.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &v1.Image{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get OpenShift Image By Name from Namespace
	image, err := cluster.ImageV1().Images().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		// Error: Getting OpenShift Image By Name & Namespace
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Image with Name: '%s' from Namespace: '%s'. Error Text: '%s'. ", name, namespace, err.Error())
		return &v1.Image{}, err
	} else {
		// Success: Getting OpenShift Image By Name & Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting OpenShift Image with Name: '%s' from Namespace: '%s'.", name, namespace)
		// Return Image
		return image, err
	}
}

// Get all OpenShift Image Streams by Namespace
func (c Openshift) GetAllImageStreams(namespace string) (*v1.ImageStreamList, error) {
	// Uses Custom OpenShift Structs
	cluster, err := imagev1.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &v1.ImageStreamList{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get All OpenShift Image Streams By Namespace
	images, err := cluster.ImageV1().ImageStreams(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		// Error: Getting All OpenShift Image Streams By Namespace
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Image Streams from Namespace: '%s'. Error Text: '%s'. ", namespace, err.Error())
		return &v1.ImageStreamList{}, err
	} else {
		// Success: Getting All OpenShift Image Streams By Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting OpenShift Image Streams from Namespace: '%s'.", namespace)
		// Return Image Streams
		return images, err
	}
}

// Get OpenShift Image Stream by Name from Namespace
func (c Openshift) GetImageStream(name string, namespace string) (*v1.ImageStream, error) {
	// Uses Custom OpenShift Structs
	cluster, err := imagev1.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &v1.ImageStream{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get OpenShift Image Stream By Name from Namespace
	image, err := cluster.ImageV1().ImageStreams(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		// Error: Getting OpenShift Image Stream By Name & Namespace
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Image Stream with Name: '%s' from Namespace: '%s'. Error Text: '%s'. ", name, namespace, err.Error())
		return &v1.ImageStream{}, err
	} else {
		// Success: Getting OpenShift Image Stream By Name & Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting OpenShift Image Stream with Name: '%s' from Namespace: '%s'.", name, namespace)
		// Return Image Stream
		return image, err
	}

}
