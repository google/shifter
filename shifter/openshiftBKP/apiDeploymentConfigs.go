package openshift

func (c *Client) DeploymentConfig(projectName string, deploymentConfigName string) (*DeploymentConfig, error) {
	req, err := c.newRequest("GET", "/apis/apps.openshift.io/v1/namespaces/"+projectName+"/deploymentconfigs/"+deploymentConfigName, nil)
	if err != nil {
		return &DeploymentConfig{}, err
	}

	deploymentConfig := &DeploymentConfig{}

	_, err = c.do(req, &deploymentConfig)
	return deploymentConfig, err
}

func (c *Client) ProjectDeploymentConfigs(projectName string) (*DeploymentConfigs, error) {
	req, err := c.newRequest("GET", "/apis/apps.openshift.io/v1/namespaces/"+projectName+"/deploymentconfigs", nil)
	if err != nil {
		return &DeploymentConfigs{}, err
	}

	deploymentConfigs := &DeploymentConfigs{}

	_, err = c.do(req, &deploymentConfigs)
	return deploymentConfigs, err
}

func (c *Client) DeploymentConfigs() (*DeploymentConfigs, error) {
	req, err := c.newRequest("GET", "/apis/apps.openshift.io/v1/deploymentconfigs", nil)
	if err != nil {
		return &DeploymentConfigs{}, err
	}

	deploymentConfigs := &DeploymentConfigs{}

	_, err = c.do(req, &deploymentConfigs)
	return deploymentConfigs, err
}
