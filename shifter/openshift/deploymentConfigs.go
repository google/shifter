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
	v1 "github.com/openshift/api/apps/v1"
	appsv1 "github.com/openshift/client-go/apps/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c Openshift) getDeploymentConfigs(namespace string) *v1.DeploymentConfigList {
	app, _ := appsv1.NewForConfig(c.clusterClient())
	depCfgLst, _ := app.AppsV1().DeploymentConfigs(namespace).List(context.TODO(), metav1.ListOptions{})

	return depCfgLst
}

func (c Openshift) getDeploymentConfig(name string, namespace string) *v1.DeploymentConfig {
	app, _ := appsv1.NewForConfig(c.clusterClient())
	depCfg, _ := app.AppsV1().DeploymentConfigs(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	return depCfg
}
