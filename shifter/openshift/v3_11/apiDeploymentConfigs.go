package v3_11

import osNativeProject "github.com/openshift/api/apps/v1"

func (a *DeploymentConfigs) Get() (*osNativeProject.DeploymentConfigList, error) {
	req, err := a.Client.NewRequest("GET", "/apis/apps.openshift.io/v1/deploymentconfigs", nil)
	if err != nil {
		return &osNativeProject.DeploymentConfigList{}, err
	}

	deploymentConfigs := &osNativeProject.DeploymentConfigList{}

	_, err = a.Client.Do(req, &deploymentConfigs)
	return deploymentConfigs, err
}

func (a *DeploymentConfigs) GetByProject(projectName string) (*osNativeProject.DeploymentConfigList, error) {
	req, err := a.Client.NewRequest("GET", "/apis/apps.openshift.io/v1/namespaces/"+projectName+"/deploymentconfigs", nil)
	if err != nil {
		return &osNativeProject.DeploymentConfigList{}, err
	}

	deploymentConfigs := &osNativeProject.DeploymentConfigList{}

	_, err = a.Client.Do(req, &deploymentConfigs)
	return deploymentConfigs, err
}
