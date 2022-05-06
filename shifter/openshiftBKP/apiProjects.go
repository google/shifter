package openshift

func (c *Client) Project(projectName string) (*Project, error) {
	req, err := c.newRequest("GET", "/apis/project.openshift.io/v1/projects/"+projectName, nil)
	if err != nil {
		return &Project{}, err
	}

	project := &Project{}

	_, err = c.do(req, &project)
	return project, err
}

func (c *Client) Projects() (*Projects, error) {
	req, err := c.newRequest("GET", "/apis/project.openshift.io/v1/projects", nil)
	if err != nil {
		return &Projects{}, err
	}

	projects := &Projects{}

	_, err = c.do(req, &projects)
	return projects, err
}
