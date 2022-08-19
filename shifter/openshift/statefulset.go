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

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func (c Openshift) GetAllStatefulSets(namespace string) (*v1.StatefulSetList, error) {
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		lib.CLog("error", "Unable to connect to cluster", err)
		return &v1.StatefulSetList{}, err
	}

	object, err := cluster.AppsV1().StatefulSets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		lib.CLog("error", "Getting StatefulSets from Namespace: "+namespace, err)
		return &v1.StatefulSetList{}, err
	}
	lib.CLog("debug", "Getting StatefulSets from Namespace: "+namespace)
	return object, err
}

func (c Openshift) GetStatefulSet(name string, namespace string) (*v1.StatefulSet, error) {
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		lib.CLog("error", "Unable to connect to cluster", err)
		return &v1.StatefulSet{}, err
	}

	object, err := cluster.AppsV1().StatefulSets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		lib.CLog("error", "Getting StatefulSet with Name: "+name+" from Namespace: "+namespace, err)
		return &v1.StatefulSet{}, err
	}
	lib.CLog("info", "Getting StatefulSet with Name: "+name+" from Namespace: "+namespace)
	return object, err
}
