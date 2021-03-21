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

package inputs

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	kyaml "sigs.k8s.io/yaml"
)

type Template struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
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
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		Required    bool   `yaml:"required"`
		Value       string `yaml:"value"`
	}
}

type kube struct {
	Parameters []struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		Required    bool   `yaml:"required"`
		Value       string `yaml:"value"`
	}
	Objects []struct {
		ApiVersion string                 `yaml:"apiVersion"`
		Kind       string                 `yaml:"kind"`
		Metadata   map[string]interface{} //metadata has a unkown structure so we use a generic interface
		Spec       map[string]interface{} //specs are dependent on the kind so we use a generic interface
		Data       map[string]interface{} //specs are dependent on the kind so we use a generic interface
	}
}

func readYaml(file string) Template {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	template := Template{}
	err = yaml.Unmarshal(yamlFile, &template)
	if err != nil {
		fmt.Println(err)
	}
	return template
}

func parse(t Template) kube {
	var k8s kube

	//iterate over the objects and modify them as needed
	for _, o := range t.Objects {
		switch o.Kind {
		case "DeploymentConfig":
			o.Kind = "Deployment"
			o.ApiVersion = "apps/v1"
			y, err := yaml.Marshal(o.Spec)
			if err != nil {
				log.Fatal(err)
			}
			j, err := kyaml.YAMLToJSON(y)
			if err != nil {
				log.Fatal(err)
			}
			//fmt.Println(string(j))
			walk(j)

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

func walk(input []byte) {
	m := map[string]interface{}{}

	err := json.Unmarshal(input, &m)
	if err != nil {
		log.Fatal(err)
	}

	for key, val := range m {
		switch cval := val.(type) {
		default:
			fmt.Println(key, ":", cval)
		}
		fmt.Println(key, val)
		//fmt.Println(val.Type)
	}

}

func TemplateConvert(input  string) []byte {
	t := readYaml(input)
	k, err := yaml.Marshal(parse(t))
			if err != nil {
				log.Println(err)
			}

			return k
}