package api

import (
	osNativeDC "github.com/openshift/api/apps/v1"
	osNativeProject "github.com/openshift/api/project/v1"
)

type ConvertItem struct {
	Namespace        *osNativeProject.Project     `json:"namespace"`
	DeploymentConfig *osNativeDC.DeploymentConfig `json:"deploymentConfig"`
	// Options * ConvertOptions `json:"options"`
}
