package api

import (
	osNative "github.com/openshift/api/apps/v1"
)

type SOSDeploymentConfigs struct {
	Shifter           Shifter                       `json:"shifter"`
	DeploymentConfigs osNative.DeploymentConfigList `json:"deploymentConfigs"`
}
