// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package processor

import (
	"encoding/json"
	"log"
	"shifter/lib"

	osroutev1 "github.com/openshift/api/route/v1"
	"golang.org/x/exp/maps"
	io "istio.io/api/networking/v1beta1"
	v1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (p Proc) Route(input []byte, flags map[string]string) lib.K8sobject {
	var route osroutev1.Route
	var processed lib.K8sobject

	err := json.Unmarshal(input, &route)
	if err != nil {
		lib.CLog("error", "Unable to parse input data for kind Route", err)
	}

	if flags["istio"] == "true" {
		if flags["create-istio-gateway"] == "Y" {
			processed = createIstioIngressGateway(route, flags)
		}

		processed = convertRouteToIstioVirtualService(route, flags)
		return processed
	} else {
		processed = convertRouteToIngress(route, flags)
		return processed
	}
}

func createIstioIngressGateway(OSRoute osroutev1.Route, flags map[string]string) lib.K8sobject {
	gw := &v1beta1.Gateway{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "networking.istio.io/v1beta1",
			Kind:       "Gateway",
		},
		ObjectMeta: OSRoute.ObjectMeta,
		Spec:       io.Gateway{},
	}

	var k lib.K8sobject
	k.Kind = gw.TypeMeta.Kind
	k.Object = gw

	return k
}

func convertRouteToIstioVirtualService(OSRoute osroutev1.Route, flags map[string]string) lib.K8sobject {
	flagIstioGateway := flags["istio-gateway"]

	vs := &v1beta1.VirtualService{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "networking.istio.io/v1beta1",
			Kind:       "VirtualService",
		},
		ObjectMeta: OSRoute.ObjectMeta,
		Spec:       io.VirtualService{},
	}

	var (
		vsSpec            io.VirtualService
		vsHTTPRoute       io.HTTPRoute
		vsHTTPMatchReq    io.HTTPMatchRequest
		vsHTTPRouteDest   io.HTTPRouteDestination
		vsDestination     io.Destination
		vsDestStringMatch io.StringMatch
		vsURIMatch        io.StringMatch_Prefix
	)

	vsSpec.Hosts = append(vsSpec.Hosts, OSRoute.Spec.Host)
	vsSpec.Gateways = append(vsSpec.Gateways, flagIstioGateway)
	// build the route
	vsHTTPRoute.Name = OSRoute.ObjectMeta.Name

	vsURIMatch.Prefix = "/"
	vsDestStringMatch.MatchType = &vsURIMatch
	vsHTTPMatchReq.Uri = &vsDestStringMatch

	vsDestination.Host = OSRoute.Spec.To.Name
	vsHTTPRouteDest.Destination = &vsDestination

	vsHTTPRoute.Match = append(vsHTTPRoute.Match, &vsHTTPMatchReq)
	vsHTTPRoute.Route = append(vsHTTPRoute.Route, &vsHTTPRouteDest)
	vsSpec.Http = append(vsSpec.Http, &vsHTTPRoute)

	vs.Spec = vsSpec

	var k lib.K8sobject
	k.Kind = vs.TypeMeta.Kind
	k.Object = vs

	return k
}

func convertRouteToIngress(OSRoute osroutev1.Route, flags map[string]string) lib.K8sobject {

	flagIngressFacing := flags["ingress-facing"]

	ingress := &v1.Ingress{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "networking.k8s.io/v1",
			Kind:       "Ingress",
		},
		ObjectMeta: OSRoute.ObjectMeta,
		Spec:       v1.IngressSpec{},
	}

	/*
		Openshift routes can take different forms which need to be handled by different types of
		ingress resources.
	*/
	var (
		ingressSpec           v1.IngressSpec
		ingressRule           v1.IngressRule
		ingressRuleValue      v1.IngressRuleValue
		httpIngressRuleValue  v1.HTTPIngressRuleValue
		ingressPath           v1.HTTPIngressPath
		ingressServiceBackend v1.IngressServiceBackend
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
	//set the path type
	var pathType v1.PathType
	pathType = "ImplementationSpecific"
	ingressPath.PathType = &pathType

	//Check if a target port has been specified
	if OSRoute.Spec.Port != nil {
		//Check if a port name has been provided otherwise use the default
		if OSRoute.Spec.Port.TargetPort.IntValue() == 0 && OSRoute.Spec.Port.TargetPort.String() != "" {
			ingressServiceBackend.Port.Name = OSRoute.Spec.Port.TargetPort.String()
		} else if OSRoute.Spec.Port.TargetPort.IntValue() != 0 {
			ingressServiceBackend.Port.Number = int32(OSRoute.Spec.Port.TargetPort.IntValue())
		}
	}
	ingressServiceBackend.Name = OSRoute.Spec.To.Name

	// build up the ingress spec
	//Add the backend service
	ingressPath.Backend.Service = &ingressServiceBackend
	httpIngressRuleValue.Paths = append(httpIngressRuleValue.Paths, ingressPath)
	ingressRuleValue.HTTP = &httpIngressRuleValue
	ingressRule.IngressRuleValue = ingressRuleValue
	ingressSpec.Rules = append(ingressSpec.Rules, ingressRule)
	ingress.Spec = ingressSpec

	// Process the flags
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

	var k lib.K8sobject
	k.Kind = ingress.TypeMeta.Kind
	k.Object = ingress

	return k
}
