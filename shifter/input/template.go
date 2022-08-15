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
	"bytes"

	gyaml "github.com/ghodss/yaml"

	//"io/ioutil"
	"log"
	"os"
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
		// Error: Unmarshalling JSON Data
		log.Printf("🧰 ❌ ERROR: Unable to Parse Input Data")
		return objects, parameters, err
	}
	return parse(template, flags)
}

func parse(template OSTemplate, flags map[string]string) (objects []lib.K8sobject, parameters []lib.OSTemplateParams, err error) {
	var k8s []lib.K8sobject
	var params []lib.OSTemplateParams

	//tplname := template.Metadata.Name

	//iterate over the objects inside the template
	for _, o := range template.Objects {
		y, _ := yaml.Marshal(o)

		jsonBody, err := gyaml.YAMLToJSON(y)
		if err != nil {
			log.Println(err)
			// TODO - ERROR HANDLING
			os.Exit(1)
		}
		// Log Opbject Conversion
		log.Printf("🧰 🚀 INFO: Converting OpenShift object of type: '%s'", o.Kind)
		processedDocument, err := processor.Processor(jsonBody, o.Kind, flags)
		if err != nil {
			// Error: Unable to Create Shifter 'TEMPLATE' Processor
			log.Printf("🧰 ❌ ERROR: Creating Shifter 'TEMPLATE' Processor.")
			return k8s, params, err
		} else {
			// Succes: Creating Shifter 'TEMPLATE' Processor
			log.Printf("🧰 ✅ SUCCESS: Shifter 'TEMPLATE' Processor Successufly Created.")
		}
		for _, v := range processedDocument {
			if v.Kind != nil {
				k8s = append(k8s, v)
			}

		}
	}

	// get the parameters from the template and store in a slice array
	for _, y := range template.Parameters {
		params = append(params, y)
	}

	// return the converted resources and parameterized values
	return k8s, params, nil
}
