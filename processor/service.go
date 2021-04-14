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
)

func convertServiceToService(OSService apiv1.Service, flags map[string]string) apiv1.Service {
	service := &apiv1.Service{
		TypeMeta:   OSService.TypeMeta,
		ObjectMeta: OSService.ObjectMeta,
		Spec:       OSService.Spec,
	}

	return *service
}
