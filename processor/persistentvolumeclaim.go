package processor

import (
	apiv1 "k8s.io/api/core/v1"
)

func convertPvcToPvc(OSPersistentVolumeClaim apiv1.PersistentVolumeClaim) apiv1.PersistentVolumeClaim {

	return OSPersistentVolumeClaim
}
