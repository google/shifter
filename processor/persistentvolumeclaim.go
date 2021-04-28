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
