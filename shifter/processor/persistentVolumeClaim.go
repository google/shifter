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
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func convertPvcToPvc(OSPersistentVolumeClaim apiv1.PersistentVolumeClaim, flags map[string]string) apiv1.PersistentVolumeClaim {

	pvc := &apiv1.PersistentVolumeClaim{
		TypeMeta: metav1.TypeMeta{
			Kind:       "PersistentVolumeClaim",
			APIVersion: "v1",
		},
		ObjectMeta: OSPersistentVolumeClaim.ObjectMeta,
		Spec:       apiv1.PersistentVolumeClaimSpec{},
	}

	var spec apiv1.PersistentVolumeClaimSpec
	spec = OSPersistentVolumeClaim.Spec
	pvc.Spec = spec

	return *pvc
}
