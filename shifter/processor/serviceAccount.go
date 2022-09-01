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
	"shifter/lib"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (p Proc) ServiceAccount(input []byte, flags map[string]string) lib.K8sobject {
	var object apiv1.ServiceAccount
	err := json.Unmarshal(input, &object)
	if err != nil {
		lib.CLog("error", "Unable to parse input data for kind: ServiceAccount", err)
	}

	sa := &apiv1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ServiceAccount",
			APIVersion: "v1",
		},
		ObjectMeta:                   object.ObjectMeta,
		Secrets:                      object.Secrets,
		ImagePullSecrets:             object.ImagePullSecrets,
		AutomountServiceAccountToken: object.AutomountServiceAccountToken,
	}
	var k lib.K8sobject
	k.Kind = sa.TypeMeta.Kind
	k.Object = sa

	return k
}
