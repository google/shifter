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

package input

import (
	"fmt"
	"io/ioutil"
	"shifter/lib"
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

func readYaml(file string) OSTemplate {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	template := OSTemplate{}
	err = yaml.Unmarshal(yamlFile, &template)
	if err != nil {
		fmt.Println(err)
	}
	return template
}

func parse(t OSTemplate) lib.Kube {
	var k8s lib.Kube

	//iterate over the objects and modify them as needed
	for _, o := range t.Objects {
		switch o.Kind {
		case "DeploymentConfig":
			//fmt.Println(o)

			//dc, _ := yaml.Marshal(o)
			//processor.DeploymentConfig(dc)

			o.Kind = "Deployment"
			o.ApiVersion = "apps/v1"
			k8s.Objects = append(k8s.Objects, o)

		case "ImageStream":
		case "Route":
		case "BuildConfig":
		case "Build":
		default:
			k8s.Objects = append(k8s.Objects, o)
		}
	}

	for _, y := range t.Parameters {
		k8s.Parameters = append(k8s.Parameters, y)
	}
	return k8s
}

func Template(input string) lib.Kube {
	t := readYaml(input)
	return parse(t)
}
