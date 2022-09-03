// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package input

import (
	"bytes"
	gyaml "github.com/ghodss/yaml"
	"shifter/lib"
	"shifter/processor"
	"sigs.k8s.io/yaml"
)

type OSTemplate struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Message    string `yaml:"message,omitempty"`
	Metadata   struct {
		CreationTimestamp string                 `yaml:"creationTimestamp"`
		Name              string                 `yaml:"name"`
		Annotations       map[string]interface{} //annotations have a unkown structure so we use a generic interface
	}
	Objects []struct {
		ApiVersion string                 `yaml:"apiVersion"`
		Kind       string                 `yaml:"kind"`
		Metadata   map[string]interface{} //metadata has a unkown structure so we use a generic interface
		Spec       map[string]interface{} //specs are dependent on the kind so we use a generic interface
		Data       map[string]interface{} //specs are dependent on the kind so we use a generic interface
	}
	Parameters []struct {
		Name string `yaml:"name"`
		//	DisplayName string `yaml:"displayName,omitempty"`
		Description string `yaml:"description,omitempty"`
		Required    bool   `yaml:"required,omitempty"`
		Value       string `yaml:"value,omitempty"`
		//	Generate string `yaml:"generate,omitempty"`
	}
}

type OSTemplateParams struct {
	Parameters []struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description,omitempty"`
		Required    bool   `yaml:"required,omitempty"`
		Value       string `yaml:"value,omitempty"`
	}
}

func Template(input bytes.Buffer, flags map[string]string) (objects []lib.K8sobject, parameters []lib.OSTemplateParams, err error) {
	template := OSTemplate{}
	err = yaml.Unmarshal(input.Bytes(), &template)
	if err != nil {
		lib.CLog("error", "Unable to parse template input data", err)
		return objects, parameters, err
	}
	return parse(template, flags)
}

func parse(template OSTemplate, flags map[string]string) (objects []lib.K8sobject, parameters []lib.OSTemplateParams, err error) {
	var k8s []lib.K8sobject
	var params []lib.OSTemplateParams

	for _, o := range template.Objects {
		y, _ := yaml.Marshal(o)

		jsonBody, err := gyaml.YAMLToJSON(y)
		if err != nil {
			lib.CLog("error", "Unable to convert yaml to json", err)
			return k8s, params, err
		}
		// Log Opbject Conversion
		lib.CLog("info", "Converting OpenShift object of type: "+o.Kind)
		processedDocument, err := processor.Processor(jsonBody, o.Kind, flags)
		if err != nil {
			lib.CLog("error", "Creating shifter processor", err)
			return k8s, params, err
		}

		for _, v := range processedDocument {
			if v.Kind != nil {
				k8s = append(k8s, v)
			}
		}
	}

	for _, y := range template.Parameters {
		params = append(params, y)
	}

	return k8s, params, nil
}
