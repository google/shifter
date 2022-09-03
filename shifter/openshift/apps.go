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

	osappsv1 "github.com/openshift/api/apps/v1"
	os "github.com/openshift/client-go/apps/clientset/versioned"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func (c Openshift) GetAllDeployments(namespace string) (*v1.DeploymentList, error) {
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		lib.CLog("error", "Unable to connect to cluster", err)
		return &v1.DeploymentList{}, err
	}

	deploymentList, err := cluster.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		lib.CLog("error", "Getting Deployments from Namespace: "+namespace, err)
		return &v1.DeploymentList{}, err
	}
	lib.CLog("debug", "Getting Deployments from Namespace: "+namespace)
	return deploymentList, nil
}

func (c Openshift) GetDeployment(name string, namespace string) (*v1.Deployment, error) {
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		lib.CLog("error", "Unable to connect to cluster", err)
		return &v1.Deployment{}, err
	}

	deployment, err := cluster.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		lib.CLog("error", "Getting Deployment with Name: "+name+" from Namespace: "+namespace, err)
		return &v1.Deployment{}, err
	}

	lib.CLog("debug", "Getting Deployment with Name: "+name+" from Namespace: "+namespace)
	return deployment, nil
}

func (c Openshift) GetAllDeploymentConfigs(namespace string) (*osappsv1.DeploymentConfigList, error) {
	cluster, err := os.NewForConfig(c.clusterClient())
	if err != nil {
		lib.CLog("error", "Unable to connect to cluster", err)
		return &osappsv1.DeploymentConfigList{}, err
	}

	depCfgLst, err := cluster.AppsV1().DeploymentConfigs(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		lib.CLog("error", "Getting All DeploymentConfigs from Namespace: "+namespace, err)
		return &osappsv1.DeploymentConfigList{}, err
	}

	lib.CLog("debug", "Getting DeploymentConfigs from Namespace: "+namespace)
	return depCfgLst, nil
}

func (c Openshift) GetDeploymentConfig(namespace string, name string) (*osappsv1.DeploymentConfig, error) {
	cluster, err := os.NewForConfig(c.clusterClient())
	if err != nil {
		lib.CLog("error", "Unable to connect to cluster", err)
		return &osappsv1.DeploymentConfig{}, err
	}

	depCfg, err := cluster.AppsV1().DeploymentConfigs(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		lib.CLog("error", "Getting DeploymentConfig with Name: "+name+" from Namespace: "+namespace, err)
		return &osappsv1.DeploymentConfig{}, err
	}
	lib.CLog("debug", "Getting DeploymentConfig with Name: "+name+" from Namespace: "+namespace)
	return depCfg, nil
}
