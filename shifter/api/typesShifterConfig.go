package api

import (
	osNativeDC "github.com/openshift/api/apps/v1"
	osNativeProject "github.com/openshift/api/project/v1"
)

type Convert struct {
	Shifter *ShifterConfig `json:"shifter"`
	Items   []*ConvertItem `json:"items"`
}

type ConvertItem struct {
	Namespace        *osNativeProject.Project     `json:"namespace"`
	DeploymentConfig *osNativeDC.DeploymentConfig `json:"deploymentConfig"`
	// Options * ConvertOptions `json:"options"`
}

type ShifterConfig struct {
	ClusterConfig *ClusterConfig `json:"clusterConfig"`
}

type ClusterConfig struct {
	ConnectionName string `json:"connectionName"`
	BaseUrl        string `json:"baseUrl"`
	BearerToken    string `json:"bearerToken"`
}
