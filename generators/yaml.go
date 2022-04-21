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

package generator

import (
	"bufio"
	"bytes"
	"fmt"
	k8sjson "k8s.io/apimachinery/pkg/runtime/serializer/json"
	"log"
	"shifter/lib"
	"strconv"
)

func (generator *Generator) yaml(name string, objects []lib.K8sobject) []lib.Converted {
	var converted []lib.Converted

	for k, v := range objects {
		no := strconv.Itoa(k)
		kind := fmt.Sprintf("%v", v.Kind)

		log.Println("Generating file", no, " - ", name)
		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)

		yaml := k8sjson.NewYAMLSerializer(k8sjson.DefaultMetaFactory, nil, nil)
		err := yaml.Encode(v.Object, writer)
		if err != nil {
			fmt.Println(err)
		}
		writer.Flush()

		var result lib.Converted
		result.Name = kind + ".yaml"
		result.Path = "/"
		result.Payload = *buff

		converted = append(converted, result)
	}

	return converted
}
