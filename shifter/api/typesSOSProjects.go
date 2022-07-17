package api

import osNativeProject "github.com/openshift/api/project/v1"

type SOSProjects struct {
	Shifter  Shifter                     `json:"shifter"`
	Projects osNativeProject.ProjectList `json:"projects"`
}
