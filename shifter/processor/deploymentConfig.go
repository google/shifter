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
	//"fmt"
	"encoding/json"
	"log"
	"shifter/lib"
	"strings"

	osappsv1 "github.com/openshift/api/apps/v1"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (p Proc) DeploymentConfig(input []byte, flags map[string]string) lib.K8sobject {
	var object osappsv1.DeploymentConfig
	err := json.Unmarshal(input, &object)
	if err != nil {
		lib.CLog("error", "Unable to parse input data for kind: DeploymentConfig", err)
	}

	flagImageRepo := flags["image-repo"]
	//fmt.Println(OSDeploymentConfig)
	// Create the body of our kubernetes deployment
	deployment := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: object.ObjectMeta,
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(object.Spec.Replicas),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: object.Spec.Template.ObjectMeta,
				Spec: apiv1.PodSpec{
					SecurityContext: &apiv1.PodSecurityContext{},
					Containers:      []apiv1.Container{},
					Volumes:         object.Spec.Template.Spec.Volumes,
				},
			},
		},
	}

	// Add the selectors to our matchlabels section in deployment.spec.selector.matchlabels
	for k, v := range object.Spec.Selector {
		deployment.Spec.Selector.MatchLabels[k] = v
	}

	// Add Volumes

	// Add Spec
	deployment.Spec.Template.Spec = object.Spec.Template.Spec

	// Add security context
	deployment.Spec.Template.Spec.SecurityContext = object.Spec.Template.Spec.SecurityContext

	// Add containers
	deployment.Spec.Template.Spec.Containers = object.Spec.Template.Spec.Containers
	for i, containers := range deployment.Spec.Template.Spec.Containers {
		if flagImageRepo != "" {
			newImg := strings.Split(containers.Image, "/")
			n := string(newImg[len(newImg)-1])
			n = flagImageRepo + n
			log.Printf("🧰 🔄 INFO: Modifying image registry source from '%s' to '%s'", containers.Image, n)
			deployment.Spec.Template.Spec.Containers[i].Image = n

		}
	}

	var k lib.K8sobject
	k.Kind = deployment.TypeMeta.Kind
	k.Object = deployment

	return k
}
