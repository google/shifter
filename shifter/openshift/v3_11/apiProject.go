package v3_11

import osNativeProject "github.com/openshift/api/project/v1"

func (a *Project) Get(projectName string) (*osNativeProject.Project, error) {
	req, err := a.Client.NewRequest("GET", "/apis/project.openshift.io/v1/projects/"+projectName, nil)
	if err != nil {
		return &osNativeProject.Project{}, err
	}

	project := &osNativeProject.Project{}

	_, err = a.Client.Do(req, &project)
	return project, err
}
