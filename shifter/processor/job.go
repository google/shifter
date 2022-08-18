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
	v1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"shifter/lib"
)

func (p Proc) Job(input []byte, flags map[string]string) lib.K8sobject {
	var object v1.Job
	err := json.Unmarshal(input, &object)
	if err != nil {
		lib.CLog("error", "Unable to parse input data for kind: Job", err)
	}

	job := &v1.Job{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Job",
			APIVersion: "batch/v1",
		},
		ObjectMeta: object.ObjectMeta,
		Spec:       object.Spec,
	}

	var k lib.K8sobject
	k.Kind = job.TypeMeta.Kind
	k.Object = job

	return k
}
