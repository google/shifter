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

package processor

import (
	"reflect"
	"shifter/lib"
)

type Proc struct{}

func int32Ptr(i int32) *int32 { return &i }
func int64Ptr(i int64) *int64 { return &i }

func Processor(input []byte, kind interface{}, flags map[string]string) ([]lib.K8sobject, error) {
	// Use our K8sobject which is a generic json interface for kubernetes objects
	var processed []lib.K8sobject

	lib.CLog("debug", "Converting "+kind.(string))

	p := Proc{}

	objects := make([]reflect.Value, 2)
	objects[0] = reflect.ValueOf(input)
	objects[1] = reflect.ValueOf(flags)

	if reflect.ValueOf(p).MethodByName(kind.(string)).IsValid() {
		m := reflect.ValueOf(p).MethodByName(kind.(string)).Call(objects)
		result := m[0].Interface().(lib.K8sobject)
		processed = append(processed, result)
	} else {
		lib.CLog("warn", "Processor doesn't exist for resource type "+kind.(string))
	}

	return processed, nil
}
