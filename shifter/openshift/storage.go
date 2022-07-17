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

func (c Openshift) GetAllPV(namespace string) (*v1.PersistentVolumeList, error) {
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		log.Println(err)
		return &v1.PersistentVolumeList{}, err
	}

	buildList, err := cluster.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Println(err)
		return &v1.PersistentVolumeList{}, err
	}

	return buildList, nil

}

func (c Openshift) GetPV(name string, namespace string) (*v1.PersistentVolume, error) {
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		log.Println(err)
		return &v1.PersistentVolume{}, err
	}

	build, err := cluster.CoreV1().PersistentVolumes().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		log.Println(err)
		return &v1.PersistentVolume{}, err
	}

	return build, nil
}

func (c Openshift) GetAllPVC(namespace string) (*v1.PersistentVolumeClaimList, error) {
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		log.Println(err)
		return &v1.PersistentVolumeClaimList{}, err
	}

	pvcList, err := cluster.CoreV1().PersistentVolumeClaims(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Println(err)
		return &v1.PersistentVolumeClaimList{}, err
	}

	return pvcList, nil

}

func (c Openshift) GetPVC(name string, namespace string) (*v1.PersistentVolumeClaim, error) {
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		log.Println(err)
		return &v1.PersistentVolumeClaim{}, err
	}

	pvc, err := cluster.CoreV1().PersistentVolumeClaims(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		log.Println(err)
		return &v1.PersistentVolumeClaim{}, err
	}

	return pvc, nil
}
