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

package processor

import (
	"fmt"
	osappsv1 "github.com/openshift/api/apps/v1"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

func convertDeploymentConfigToDeployment(OSDeploymentConfig osappsv1.DeploymentConfig, flags map[string]string) appsv1.Deployment {

	var flagImageRepo string
	flagImageRepo = flags["image-repo"]

	// Create the body of our kubernetes deployment
	deployment := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: OSDeploymentConfig.ObjectMeta,
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(OSDeploymentConfig.Spec.Replicas),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: OSDeploymentConfig.Spec.Template.ObjectMeta,
				Spec: apiv1.PodSpec{
					SecurityContext: &apiv1.PodSecurityContext{},
					Containers:      []apiv1.Container{},
					Volumes:         OSDeploymentConfig.Spec.Template.Spec.Volumes,
				},
			},
		},
	}

	// Add the selectors to our matchlabels section in deployment.spec.selector.matchlabels
	for k, v := range OSDeploymentConfig.Spec.Selector {
		deployment.Spec.Selector.MatchLabels[k] = v
	}

	// Add Volumes

	// Add Spec
	deployment.Spec.Template.Spec = OSDeploymentConfig.Spec.Template.Spec

	// Add security context
	deployment.Spec.Template.Spec.SecurityContext = OSDeploymentConfig.Spec.Template.Spec.SecurityContext

	// Add containers
	deployment.Spec.Template.Spec.Containers = OSDeploymentConfig.Spec.Template.Spec.Containers
	for i, containers := range deployment.Spec.Template.Spec.Containers {
		if flagImageRepo != "" {
			var newImg string
			newImg = strings.Split(containers.Image, "/")[1]
			fmt.Println(newImg[1])
			newImg = flagImageRepo + newImg
			fmt.Println(newImg)
			deployment.Spec.Template.Spec.Containers[i].Image = newImg

		}
	}

	// Return a full kubernetes structure, this needs to be marshalled into a usable yaml
	return *deployment
}
