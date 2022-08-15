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
	"shifter/lib"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func convertDeploymentToDeployment(OSDeployment v1.Deployment, flags map[string]string) lib.K8sobject {
	deployment := &v1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       DEPLOYMENT,
			APIVersion: "apps/v1",
		},
		ObjectMeta: OSDeployment.ObjectMeta,
		Spec:       OSDeployment.Spec,
	}

	var k lib.K8sobject
	k.Kind = DEPLOYMENT
	k.Object = deployment

	return k
}

func convertDaemonSetToDaemonSet(OSDaemonSet v1.DaemonSet, flags map[string]string) lib.K8sobject {
	daemonset := &v1.DaemonSet{
		TypeMeta: metav1.TypeMeta{
			Kind:       DAEMONSET,
			APIVersion: "apps/v1",
		},
		ObjectMeta: OSDaemonSet.ObjectMeta,
		Spec:       OSDaemonSet.Spec,
	}

	var k lib.K8sobject
	k.Kind = DAEMONSET
	k.Object = daemonset

	return k
}

func convertStatefulSetToStatefulSet(OSStatefulSet v1.StatefulSet, flags map[string]string) lib.K8sobject {
	statefulset := &v1.StatefulSet{
		TypeMeta: metav1.TypeMeta{
			Kind:       STATEFULSET,
			APIVersion: "apps/v1",
		},
		ObjectMeta: OSStatefulSet.ObjectMeta,
		Spec:       OSStatefulSet.Spec,
	}

	var k lib.K8sobject
	k.Kind = STATEFULSET
	k.Object = statefulset

	return k
}
