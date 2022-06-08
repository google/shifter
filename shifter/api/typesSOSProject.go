package api

import osNativeProject "github.com/openshift/api/project/v1"

type SOSProject struct {
	Shifter Shifter                 `json:"shifter"`
	Project osNativeProject.Project `json:"project"`
}
