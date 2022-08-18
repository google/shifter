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
	"errors"
	"fmt"
	"log"
	"shifter/lib"

	osappsv1 "github.com/openshift/api/apps/v1"
	osroutev1 "github.com/openshift/api/route/v1"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
)

const (
	SECRET                = "Secret"
	PERSISTENTVOLUME      = "PersistentVolume"
	PERSISTENTVOLUMECLAIM = "PersistentVolumeClaim"
	GATEWAY               = "Gateway"
	JOB                   = "Job"
	TEMPLATE              = "Template"
	IMAGESTREAM           = "ImageStream"
	CONFIGMAP             = "ConfigMap"
	BUILD                 = "Build"
	STATEFULSET           = "StatefulSet"
	SERVICEACCOUNT        = "ServiceAccount"
	DAEMONSET             = "DaemonSet"
	DEPLOYMENT            = "Deployment"
	DEPLOYMENTCONFIG      = "DeploymentConfig"
	SERVICE               = "Service"
	ROUTE                 = "Route"
	VIRTUALSERVICE        = "VirtualService"
	INGRESS               = "Ingress"
)

func int32Ptr(i int32) *int32 { return &i }
func int64Ptr(i int64) *int64 { return &i }

func Processor(input []byte, kind interface{}, flags map[string]string) ([]lib.K8sobject, error) {
	// Use our K8sobject which is a generic json interface for kubernetes objects
	var processed []lib.K8sobject

	switch kind {
	case DEPLOYMENTCONFIG:
		var object osappsv1.DeploymentConfig
		err := json.Unmarshal(input, &object)
		if err != nil {
			// Error: Unmarshalling JSON Data
			log.Printf("🧰 ❌ ERROR: Unable to Parse Input Data for Kind: '%s'.", kind)
			return processed, err
		}
		processed := append(processed, convertDeploymentConfigToDeployment(object, flags))
		return processed, nil // Success
		break

	case DEPLOYMENT:
		var object appsv1.Deployment
		err := json.Unmarshal(input, &object)
		if err != nil {
			// Error: Unmarshalling JSON Data
			log.Printf("🧰 ❌ ERROR: Unable to Parse Input Data for Kind: '%s'.", kind)
			return processed, err
		}
		processed := append(processed, convertDeploymentToDeployment(object, flags))
		return processed, nil // Success
		break

	case STATEFULSET:
		var object appsv1.StatefulSet
		err := json.Unmarshal(input, &object)
		if err != nil {
			// Error: Unmarshalling JSON Data
			log.Printf("🧰 ❌ ERROR: Unable to Parse Input Data for Kind: '%s'.", kind)
			return processed, err
		}
		processed := append(processed, convertStatefulSetToStatefulSet(object, flags))
		return processed, nil // Success
		break

	case DAEMONSET:
		var object appsv1.DaemonSet
		err := json.Unmarshal(input, &object)
		if err != nil {
			// Error: Unmarshalling JSON Data
			log.Printf("🧰 ❌ ERROR: Unable to Parse Input Data for Kind: '%s'.", kind)
			return processed, err
		}
		processed := append(processed, convertDaemonSetToDaemonSet(object, flags))
		return processed, nil // Success
		break

	case ROUTE:
		var route osroutev1.Route
		err := json.Unmarshal(input, &route)
		if err != nil {
			// Error: Unmarshalling JSON Data
			log.Printf("🧰 ❌ ERROR: Unable to Parse Input Data for Kind: '%s'.", kind)
			return processed, err
		}

		if flags["istio"] == "true" {
			if flags["create-istio-gateway"] == "Y" {
				processed = append(processed, createIstioIngressGateway(route, flags))
			}

			processed = append(processed, convertRouteToIstioVirtualService(route, flags))
			return processed, nil // Success
			break
		} else {
			processed := append(processed, convertRouteToIngress(route, flags))
			return processed, nil // Success
			break
		}

	case SERVICE:
		var service apiv1.Service
		err := json.Unmarshal(input, &service)
		if err != nil {
			// Error: Unmarshalling JSON Data
			log.Printf("🧰 ❌ ERROR: Unable to Parse Input Data for Kind: '%s'.", kind)
			return processed, err
		}
		processed := append(processed, convertServiceToService(service, flags))
		return processed, nil // Success
		break

	case CONFIGMAP:
		var cfgMap apiv1.ConfigMap
		err := json.Unmarshal(input, &cfgMap)
		if err != nil {
			// Error: Unmarshalling JSON Data
			log.Printf("🧰 ❌ ERROR: Unable to Parse Input Data for Kind: '%s'.", kind)
			return processed, err
		}
		processed := append(processed, convertConfigMapToConfigMap(cfgMap, flags))
		return processed, nil // Success
		break

	case SERVICEACCOUNT:
		var sa apiv1.ServiceAccount
		err := json.Unmarshal(input, &sa)
		if err != nil {
			// Error: Unmarshalling JSON Data
			log.Printf("🧰 ❌ ERROR: Unable to Parse Input Data for Kind: '%s'.", kind)
			return processed, err
		}
		processed := append(processed, convertServiceAccountToServiceAccount(sa, flags))
		return processed, nil // Success
		break
	}

	// Error Unsupported Processor Type
	return processed, errors.New(fmt.Sprintf("🧰 ❌ ERROR: Unsupported Processor Type: '%s'", kind))
}

// TODO - Remove Function
/*
func serializer(input runtime.Object) {
	fmt.Println("---")
	e := kjson.NewYAMLSerializer(kjson.DefaultMetaFactory, nil, nil)

	err := e.Encode(input, os.Stdout)
	if err != nil {
		fmt.Println(err)
	}
}*/
