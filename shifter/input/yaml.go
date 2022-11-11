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
