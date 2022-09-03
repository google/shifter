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

package generator

import (
	"reflect"
	"shifter/lib"
	"strings"
)

type Generator struct{}

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

	return result, nil
}
