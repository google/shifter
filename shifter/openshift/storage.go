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

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func (c Openshift) GetAllPV(namespace string) (*v1.PersistentVolumeList, error) {
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		lib.CLog("error", "Unable to connect to cluster", err)
		return &v1.PersistentVolumeList{}, err
	}

	pvs, err := cluster.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		lib.CLog("error", "Getting PersistentVolumes from Namespace: "+namespace, err)
		return &v1.PersistentVolumeList{}, err
	}
	lib.CLog("debug", "Getting PersistentVolumes from Namespace: "+namespace)
	return pvs, err
}

func (c Openshift) GetPV(name string, namespace string) (*v1.PersistentVolume, error) {
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		lib.CLog("error", "Unable to connect to cluster", err)
		return &v1.PersistentVolume{}, err
	}

	pv, err := cluster.CoreV1().PersistentVolumes().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		lib.CLog("error", "Getting PersistentVolume with Name: "+name+" from Namespace: "+namespace, err)
		return &v1.PersistentVolume{}, err
	}
	lib.CLog("info", "Getting PersistentVolume with Name: "+name+" from Namespace: "+namespace)
	return pv, err
}

func (c Openshift) GetAllPVC(namespace string) (*v1.PersistentVolumeClaimList, error) {
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		lib.CLog("error", "Unable to connect to cluster", err)
		return &v1.PersistentVolumeClaimList{}, err
	}

	pvcList, err := cluster.CoreV1().PersistentVolumeClaims(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		lib.CLog("error", "Getting PersistentVolumeClaims from Namespace: "+namespace, err)
		return &v1.PersistentVolumeClaimList{}, err
	}
	lib.CLog("debug", "Getting PersistentVolumeClaims from Namespace: "+namespace)
	return pvcList, err
}

func (c Openshift) GetPVC(name string, namespace string) (*v1.PersistentVolumeClaim, error) {
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		lib.CLog("error", "Unable to connect to cluster", err)
		return &v1.PersistentVolumeClaim{}, err
	}

	pvc, err := cluster.CoreV1().PersistentVolumeClaims(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		lib.CLog("error", "Getting PersistentVolumeClaim with Name: "+name+" from Namespace: "+namespace, err)
		return &v1.PersistentVolumeClaim{}, err
	}
	lib.CLog("info", "Getting PersistentVolumeClaim with Name: "+name+" from Namespace: "+namespace)
	return pvc, err
}
