package v3_11

import osNativeProject "github.com/openshift/api/project/v1"


func (a *Projects) Get() (*osNativeProject.ProjectList, error) {
	req, err := a.Client.NewRequest("GET", "/apis/project.openshift.io/v1/projects", nil)
	if err != nil {
		return &osNativeProject.ProjectList{}, err
	}
	projects := &osNativeProject.ProjectList{}
	_, err = a.Client.Do(req, &projects)
	return projects, err
}
