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
	//"github.com/instrumenta/kubeval/kubeval"
	"reflect"
	"shifter/lib"
	"strings"
)

type Generator struct {
	/*
		OutputType string
		Name       string
		Input      struct {
			Object     []lib.K8sobject
			Parameters []lib.OSTemplateParams
		}
	*/
}

type Input struct {
	Object []lib.K8sobject
}

func NewGenerator(outputType string, args ...interface{}) ([]lib.Converted, error) {
	c := Generator{}
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}

	outputType = strings.Title(strings.ToLower(outputType))
	val := reflect.ValueOf(c).MethodByName(outputType).Call(inputs)

	var result []lib.Converted
	result = val[0].Interface().([]lib.Converted)

	// Success, New Generator Created
	return result, nil
}
