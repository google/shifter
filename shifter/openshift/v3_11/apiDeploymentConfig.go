package v3_11

import (
	"fmt"

	osNativeProject "github.com/openshift/api/apps/v1"
)

func (a *DeploymentConfig) Get(projectName string, deploymentConfigName string) (*osNativeProject.DeploymentConfig, error) {
	req, err := a.Client.NewRequest("GET", "/apis/apps.openshift.io/v1/namespaces/"+projectName+"/deploymentconfigs/"+deploymentConfigName, nil)
	if err != nil {
		fmt.Println("OH SHIT!")
		return &osNativeProject.DeploymentConfig{}, err
	}
	fmt.Println(projectName)
	fmt.Println(deploymentConfigName)
	fmt.Println(req)
	deploymentConfig := &osNativeProject.DeploymentConfig{}

	_, err = a.Client.Do(req, &deploymentConfig)
	return deploymentConfig, err
}
