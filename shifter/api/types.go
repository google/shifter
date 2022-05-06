/*
Copyright 2019 Google LLC
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import (
	osNativeProject "github.com/openshift/api/project/v1"
)

type ShifterClusterConfig struct {
	BaseUrl     string `json:"baseUrl"`
	BasePort    string `json:"basePort"`
	BearerToken string `json:"bearerToken"`
}

type Shifter struct {
	ClusterConfig ShifterClusterConfig `json:"clusterConfig"`
}

type ShifterGetOpenShiftProjects struct {
	Shifter  Shifter                     `json:"shifter"`
	Projects osNativeProject.ProjectList `json:"projects"`
}

type ShifterGetOpenShiftProject struct {
	Shifter Shifter                 `json:"shifter"`
	Project osNativeProject.Project `json:"project"`
}
