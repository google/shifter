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
	"reflect"
	"shifter/lib"
	"strings"
)

type Importer struct{}

type blah struct {
	Object     []lib.K8sobject
	Parameters []lib.OSTemplateParams
}

func Import(inputType string, args ...interface{}) []blah {
	c := Importer{}
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}

	inputType = strings.Title(strings.ToLower(inputType))
	val := reflect.ValueOf(c).MethodByName(inputType).Call(inputs)

	var result []blah
	result = val[0].Interface().([]blah)

	return result
}
