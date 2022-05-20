package v3_11

import (
	osNativeProject "github.com/openshift/api/apps/v1"
)

func (a *DeploymentConfig) Get(projectName string, deploymentConfigName string) (*osNativeProject.DeploymentConfig, error) {
	req, err := a.Client.NewRequest("GET", "/apis/apps.openshift.io/v1/namespaces/"+projectName+"/deploymentconfigs/"+deploymentConfigName, nil)
	if err != nil {
		return &osNativeProject.DeploymentConfig{}, err
	}
	deploymentConfig := &osNativeProject.DeploymentConfig{}

	_, err = a.Client.Do(req, &deploymentConfig)
	return deploymentConfig, err
}
