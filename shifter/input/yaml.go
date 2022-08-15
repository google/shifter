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

// TODO - Error Handling
func Yaml(input bytes.Buffer, flags map[string]string) ([]lib.K8sobject, error) {

	//nestedQuotedStringHack(fileName)
	f := bytes.NewReader(input.Bytes())
	d := yaml.NewDecoder(f)
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
			// Error: Unable to Marshal YAML
			log.Printf("🧰 ❌ ERROR: Unable to Marshal YAML.")
			return nil, err
		}

		jsonBody, err := gyaml.YAMLToJSON(val)
		if err != nil {
			// Error: Unable to Parse YAML to JSON
			log.Printf("🧰 ❌ ERROR: Unable to Parse YAML to JSON.")
			return nil, err
		}

		processedDocument, err := processor.Processor(jsonBody, doc["kind"], flags)
		if err != nil {
			// Error: Unable to Create Shifter 'YAML' Processor
			log.Printf("🧰 ❌ ERROR: Creating Shifter 'YAML' Processor.")
			return nil, err
		} else {
			// Succes: Creating Shifter 'YAML' Processor
			log.Printf("🧰 ✅ SUCCESS: Shifter 'YAML' Processor Successufly Created.")
		}
		for _, v := range processedDocument {
			objects = append(objects, v)
		}
	}
	// Success
	return objects, nil
}

// TODO - Remove this Function
/*
func nestedQuotedStringHack(fileName string) {
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println(err)
	}

	str1 := string(input)
	lines := strings.Split(str1, "\n")

	for i, line := range lines {
		found := strings.Contains(line, `\"`)
		if found == true {
			if strings.HasSuffix(lines[i], `'`) == false {
				lines[i] = strings.Replace(lines[i], `"`, `'`, 1)
				lines[i] = strings.TrimSuffix(lines[i], `"`)
				lines[i] = lines[i] + `'`
			}
		}
	}

	output := strings.Join(lines, "\n")

	err = ioutil.WriteFile(fileName, []byte(output), 0644)
	if err != nil {
		log.Println(err)
	}
}*/
