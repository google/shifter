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
	v1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"shifter/lib"
)

func convertJobtoJob(OSJob v1.Job, flags map[string]string) lib.K8sobject {
	job := &v1.Job{
		TypeMeta: metav1.TypeMeta{
			Kind:       JOB,
			APIVersion: "batch/v1",
		},
		ObjectMeta: OSJob.ObjectMeta,
		Spec:       OSJob.Spec,
	}

	var k lib.K8sobject
	k.Kind = JOB
	k.Object = job

	return k
}
