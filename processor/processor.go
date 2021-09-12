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
	"github.com/google/shifter/lib"
	osappsv1 "github.com/openshift/api/apps/v1"
	osroutev1 "github.com/openshift/api/route/v1"
	apiv1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	kjson "k8s.io/apimachinery/pkg/runtime/serializer/json"
	"os"
)

func int32Ptr(i int32) *int32 { return &i }
func int64Ptr(i int64) *int64 { return &i }

func Processor(input []byte, kind interface{}, flags map[string]string) lib.K8sobject {
	// Use our K8sobject which is a generic json interface for kubernetes objects
	var k lib.K8sobject

	switch kind {
	case "DeploymentConfig":
		var dc osappsv1.DeploymentConfig
		json.Unmarshal(input, &dc)
		t := convertDeploymentConfigToDeployment(dc, flags)
		var k lib.K8sobject

		k.Kind = "Deployment"
		k.Object = &t

		return k
		break

	case "Route":
		var route osroutev1.Route
		json.Unmarshal(input, &route)
		t := convertRouteToIngress(route, flags)
		var k lib.K8sobject

		k.Kind = "Ingress"
		k.Object = &t

		return k
		break

	case "Service":
		var service apiv1.Service
		json.Unmarshal(input, &service)
		t := convertServiceToService(service, flags)
		var k lib.K8sobject

		k.Kind = kind
		k.Object = &t

		return k
		break

	case "PersistentVolumeClaim":
		var pvc apiv1.PersistentVolumeClaim
		json.Unmarshal(input, &pvc)
		t := convertPvcToPvc(pvc, flags)
		var k lib.K8sobject

		k.Kind = kind
		k.Object = &t

		return k
		break

	case "ConfigMap":
		var cfgMap apiv1.ConfigMap
		json.Unmarshal(input, &cfgMap)
		t := convertConfigMapToConfigMap(cfgMap, flags)
		var k lib.K8sobject

		k.Kind = kind
		k.Object = &t

		return k
		break

	}

	return k
}

func serializer(input runtime.Object) {
	fmt.Println("---")
	e := kjson.NewYAMLSerializer(kjson.DefaultMetaFactory, nil, nil)

	err := e.Encode(input, os.Stdout)
	if err != nil {
		fmt.Println(err)
	}
}
