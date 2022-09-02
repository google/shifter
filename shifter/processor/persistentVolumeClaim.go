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
