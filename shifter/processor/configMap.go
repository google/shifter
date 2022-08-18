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
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"shifter/lib"
)

func (p Proc) ConfigMap(input []byte, flags map[string]string) lib.K8sobject {
	var object apiv1.ConfigMap
	err := json.Unmarshal(input, &object)
	if err != nil {
		lib.CLog("error", "Unable to parse input data for kind: ConfigMap", err)
	}

	cfgMap := &apiv1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: object.ObjectMeta,
		Data:       object.Data,
		BinaryData: object.BinaryData,
	}
	var k lib.K8sobject
	k.Kind = cfgMap.TypeMeta.Kind
	k.Object = cfgMap

	return k
}
