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
	yaml "gopkg.in/yaml.v3"
	"io"
	"io/ioutil"
	"log"
	"os"
	"shifter/lib"
	"shifter/processor"
	"strings"
)

type Spec struct {
	Kind string `yaml:"kind"`
}

func Yaml(input bytes.Buffer, flags map[string]string) []lib.K8sobject {

	//nestedQuotedStringHack(fileName)

	f := bytes.NewReader(input.Bytes())
	d := yaml.NewDecoder(f)
	objects := make([]lib.K8sobject, 0)

	for {
		doc := make(map[interface{}]interface{})

		err := d.Decode(&doc)
		if err != nil {
			if err != io.EOF {
				log.Println(err)
				os.Exit(1)
			}
		}

		if err == io.EOF {
			break
		}

		val, err := yaml.Marshal(doc)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		jsonBody, err := gyaml.YAMLToJSON(val)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		processedDocument := processor.Processor(jsonBody, doc["kind"], flags)
		for _, v := range processedDocument {
			objects = append(objects, v)
		}
	}
	return objects
}

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
}
