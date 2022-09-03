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
	"shifter/lib"

	v1 "github.com/openshift/api/image/v1"
	imagev1 "github.com/openshift/client-go/image/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c Openshift) GetAllImages(namespace string) (*v1.ImageList, error) {
	cluster, err := imagev1.NewForConfig(c.clusterClient())
	if err != nil {
		lib.CLog("error", "Unable to connect to cluster", err)
		return &v1.ImageList{}, err
	}

	images, err := cluster.ImageV1().Images().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		lib.CLog("error", "Getting Images from Namespace: "+namespace, err)
		return &v1.ImageList{}, err
	}
	lib.CLog("debug", "Getting Images from Namespace: "+namespace)
	return images, err
}

func (c Openshift) GetImage(name string, namespace string) (*v1.Image, error) {
	cluster, err := imagev1.NewForConfig(c.clusterClient())
	if err != nil {
		lib.CLog("error", "Unable to connect to cluster", err)
		return &v1.Image{}, err
	}

	image, err := cluster.ImageV1().Images().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		lib.CLog("error", "Getting Image with Name: "+name+" from Namespace: "+namespace, err)
		return &v1.Image{}, err
	}
	lib.CLog("info", "Getting Image with Name: "+name+" from Namespace: "+namespace)
	return image, err
}

func (c Openshift) GetAllImageStreams(namespace string) (*v1.ImageStreamList, error) {
	cluster, err := imagev1.NewForConfig(c.clusterClient())
	if err != nil {
		lib.CLog("error", "Unable to connect to cluster", err)
		return &v1.ImageStreamList{}, err
	}

	images, err := cluster.ImageV1().ImageStreams(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		lib.CLog("error", "Getting ImageStreams from Namespace: "+namespace, err)
		return &v1.ImageStreamList{}, err
	}
	lib.CLog("debug", "Getting ImageStreams from Namespace: "+namespace)
	return images, err
}

func (c Openshift) GetImageStream(name string, namespace string) (*v1.ImageStream, error) {
	cluster, err := imagev1.NewForConfig(c.clusterClient())
	if err != nil {
		lib.CLog("error", "Unable to connect to cluster", err)
		return &v1.ImageStream{}, err
	}

	image, err := cluster.ImageV1().ImageStreams(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		lib.CLog("error", "Getting ImageStream with Name: "+name+" from Namespace: "+namespace, err)
		return &v1.ImageStream{}, err
	}
	lib.CLog("info", "Getting ImageStream with Name: "+name+" from Namespace: "+namespace)
	return image, err
}
