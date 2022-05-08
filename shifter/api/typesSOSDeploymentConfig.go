package api

import (
	osNative "github.com/openshift/api/apps/v1"
)

type SOSDeploymentConfig struct {
	Shifter          Shifter                          `json:"shifter"`
	DeploymentConfig osNative.DeploymentConfig `json:"deploymentConfig"`
}
