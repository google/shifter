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
	osroutev1 "github.com/openshift/api/route/v1"
	"golang.org/x/exp/maps"
	v1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"log"
)

func convertRouteToIngress(OSRoute osroutev1.Route, flags map[string]string) v1beta1.Ingress {

	flagIngressFacing := flags["ingress-facing"]

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
		ingressSpec          v1beta1.IngressSpec
		ingressRule          v1beta1.IngressRule
		ingressRuleValue     v1beta1.IngressRuleValue
		httpIngressRuleValue v1beta1.HTTPIngressRuleValue
		ingressPath          v1beta1.HTTPIngressPath
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

	ingressPath.Backend.ServicePort = intstr.FromString(OSRoute.Spec.To.Name)
	ingressPath.Backend.ServiceName = OSRoute.Spec.To.Name

	httpIngressRuleValue.Paths = append(httpIngressRuleValue.Paths, ingressPath)
	ingressRuleValue.HTTP = &httpIngressRuleValue
	ingressRule.IngressRuleValue = ingressRuleValue
	ingressSpec.Rules = append(ingressSpec.Rules, ingressRule)
	ingress.Spec = ingressSpec

	if flagIngressFacing == "internal" {
		log.Println("Modifying ingress to internal loadbalancer")
		annotation := make(map[string]string)
		annotation["kubernetes.io/ingress.class"] = "gce-internal"
		if ingress.ObjectMeta.Annotations != nil {
			maps.Copy(ingress.ObjectMeta.Annotations, annotation)
		} else {
			ingress.ObjectMeta.Annotations = annotation
		}
	}

	return *ingress
}
