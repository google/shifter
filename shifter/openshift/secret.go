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

func (c Openshift) GetAllSecrets(namespace string) (*v1.SecretList, error) {
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		lib.CLog("error", "Unable to connect to cluster", err)
		return &v1.SecretList{}, err
	}

	secretList, err := cluster.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		lib.CLog("error", "Getting Secrets from Namespace: "+namespace, err)
		return &v1.SecretList{}, err
	}
	lib.CLog("debug", "Getting Secrets from Namespace: "+namespace)
	return secretList, err
}

func (c Openshift) GetSecret(name string, namespace string) (*v1.Secret, error) {
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		lib.CLog("error", "Unable to connect to cluster", err)
		return &v1.Secret{}, err
	}

	secret, err := cluster.CoreV1().Secrets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		lib.CLog("error", "Getting Secret with Name: "+name+" from Namespace: "+namespace, err)
		return &v1.Secret{}, err
	}
	lib.CLog("info", "Getting Secret with Name: "+name+" from Namespace: "+namespace)
	return secret, err
}
