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
	"io"
	"log"
	"shifter/lib"
	"shifter/processor"

	gyaml "github.com/ghodss/yaml"
	yaml "gopkg.in/yaml.v3"
)

type Spec struct {
	Kind string `yaml:"kind"`
}

func Yaml(input bytes.Buffer, flags map[string]string) ([]lib.K8sobject, error) {

	file := bytes.NewReader(input.Bytes())
	d := yaml.NewDecoder(file)
	objects := make([]lib.K8sobject, 0)

	for {
		doc := make(map[interface{}]interface{})

		err := d.Decode(&doc)
		if err != nil {
			if err != io.EOF {
				log.Printf("🧰 ❌ ERROR: Parsing YAML.")
				return nil, err
			}
		}

		if err == io.EOF {
			break
		}

		val, err := yaml.Marshal(doc)
		if err != nil {
			lib.CLog("error", "Unable to Marshal YAML", err)
			return nil, err
		}

		jsonBody, err := gyaml.YAMLToJSON(val)
		if err != nil {
			lib.CLog("error", "Unable to convert yaml to json.", err)
			return nil, err
		}

		processedDocument, err := processor.Processor(jsonBody, doc["kind"], flags)
		if err != nil {
			lib.CLog("error", "Creating processor", err)
			return nil, err
		}
		for _, v := range processedDocument {
			objects = append(objects, v)
		}
	}
	return objects, nil
}
