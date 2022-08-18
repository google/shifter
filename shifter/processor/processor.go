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
