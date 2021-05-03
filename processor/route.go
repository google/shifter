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
	osroutev1 "github.com/openshift/api/route/v1"
	v1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func convertRouteToIngress(OSRoute osroutev1.Route, flags map[string]string) v1beta1.Ingress {
	ingress := &v1beta1.Ingress{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "networking.k8s.io/v1beta1",
			Kind:       "Ingress",
		},
		ObjectMeta: OSRoute.ObjectMeta,
		Spec:       v1beta1.IngressSpec{},
	}

	/*
		Openshift routes can take different forms which need to be handled by different types of
		ingress resources.
	*/

	var (
		ingressSpec v1beta1.IngressSpec
		ingressRule v1beta1.IngressRule
		ingressPath v1beta1.HTTPIngressPath
	)

	//Logic to convert a route to ingress

	// Check if there is a host specified
	if OSRoute.Spec.Host != "" {
		ingressRule.Host = OSRoute.Spec.Host
	}

	//Build up the paths for the ingress resource
	if OSRoute.Spec.Path != "" {
		ingressPath.Path = OSRoute.Spec.Path
	} else {
		ingressPath.Path = "/"
	}

	var pathType v1beta1.PathType
	pathType = "Prefix"
	ingressPath.PathType = &pathType

	var backend v1beta1.IngressBackend
	backend.ServiceName = "test"
	backend.ServicePort = intstr.FromInt(8080)
	ingressPath.Backend = backend

	ingressRule.HTTP.Paths = append(ingressRule.HTTP.Paths, ingressPath)

	ingressSpec.Rules = append(ingressSpec.Rules, ingressRule)
	ingress.Spec = ingressSpec

	return *ingress
}
