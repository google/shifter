/*
copyright 2019 google llc
licensed under the apache license, version 2.0 (the "license");
you may not use this file except in compliance with the license.
you may obtain a copy of the license at
    http://www.apache.org/licenses/license-2.0
unless required by applicable law or agreed to in writing, software
distributed under the license is distributed on an "as is" basis,
without warranties or conditions of any kind, either express or implied.
see the license for the specific language governing permissions and
limitations under the license.
*/

package lib

import (
	"bytes"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

type Kube struct {
	Parameters []struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description,omitempty"`
		Required    bool   `yaml:"required,omitempty"`
		Value       string `yaml:"value,omitempty"`
	}
	Objects []struct {
		ApiVersion string                 `yaml:"apiVersion"`
		Kind       string                 `yaml:"kind"`
		Metadata   map[string]interface{} //metadata has a unkown structure so we use a generic interface
		Spec       map[string]interface{} //specs are dependent on the kind so we use a generic interface
		Data       map[string]interface{} //specs are dependent on the kind so we use a generic interface
	}
}

type OSTemplateParams struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description,omitempty"`
	Required    bool   `yaml:"required,omitempty"`
	Value       string `yaml:"value,omitempty"`
}

type K8sobject struct {
	// type of converted Openshift Resource to Kubernetes Resource
	Kind   interface{}
	Object runtime.Object
}

type Converted struct {
	Name    string
	Path    string
	Payload bytes.Buffer
}
