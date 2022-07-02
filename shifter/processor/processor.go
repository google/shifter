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
	"encoding/json"
	"fmt"
	osappsv1 "github.com/openshift/api/apps/v1"
	osroutev1 "github.com/openshift/api/route/v1"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	kjson "k8s.io/apimachinery/pkg/runtime/serializer/json"
	"os"
	"shifter/lib"
)

func int32Ptr(i int32) *int32 { return &i }
func int64Ptr(i int64) *int64 { return &i }

func Processor(input []byte, kind interface{}, flags map[string]string) []lib.K8sobject {
	// Use our K8sobject which is a generic json interface for kubernetes objects
	var processed []lib.K8sobject

	switch kind {
	case "DeploymentConfig":
		var object osappsv1.DeploymentConfig
		json.Unmarshal(input, &object)
		processed := append(processed, convertDeploymentConfigToDeployment(object, flags))
		return processed
		break

	case "Deployment":
		var object appsv1.Deployment
		json.Unmarshal(input, &object)
		processed := append(processed, convertDeploymentToDeployment(object, flags))
		return processed
		break

	case "StatefulSet":
		var object appsv1.StatefulSet
		json.Unmarshal(input, &object)
		processed := append(processed, convertStatefulSetToStatefulSet(object, flags))
		return processed
		break

	case "DaemonSet":
		var object appsv1.DaemonSet
		json.Unmarshal(input, &object)
		processed := append(processed, convertDaemonSetToDaemonSet(object, flags))
		return processed
		break

	case "Route":
		var route osroutev1.Route
		json.Unmarshal(input, &route)

		if flags["istio"] == "true" {
			if flags["create-istio-gateway"] == "Y" {
				processed = append(processed, createIstioIngressGateway(route, flags))
			}

			processed = append(processed, convertRouteToIstioVirtualService(route, flags))
			return processed
			break
		} else {
			processed := append(processed, convertRouteToIngress(route, flags))
			return processed
			break
		}

	case "Service":
		var service apiv1.Service
		json.Unmarshal(input, &service)
		processed := append(processed, convertServiceToService(service, flags))
		return processed
		break

	case "ConfigMap":
		var cfgMap apiv1.ConfigMap
		json.Unmarshal(input, &cfgMap)
		processed := append(processed, convertConfigMapToConfigMap(cfgMap, flags))
		return processed
		break

	case "ServiceAccount":
		var sa apiv1.ServiceAccount
		json.Unmarshal(input, &sa)
		processed := append(processed, convertServiceAccountToServiceAccount(sa, flags))
		return processed
		break
	}

	return processed
}

func serializer(input runtime.Object) {
	fmt.Println("---")
	e := kjson.NewYAMLSerializer(kjson.DefaultMetaFactory, nil, nil)

	err := e.Encode(input, os.Stdout)
	if err != nil {
		fmt.Println(err)
	}
}
