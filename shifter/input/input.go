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

/*
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
*/
