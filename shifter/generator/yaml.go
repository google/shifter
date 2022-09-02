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
	"bufio"
	"bytes"
	"fmt"
	"shifter/lib"

	k8sjson "k8s.io/apimachinery/pkg/runtime/serializer/json"
)

func (generator Generator) Yaml(name string, objects []lib.K8sobject) []lib.Converted {
	var converted []lib.Converted

	for _, v := range objects {
		//kind := fmt.Sprintf("%v", v.Kind)
		buff := new(bytes.Buffer)
		writer := bufio.NewWriter(buff)

		serializer := k8sjson.NewSerializerWithOptions(k8sjson.DefaultMetaFactory, nil, nil,
			k8sjson.SerializerOptions{
				Yaml:   true,
				Pretty: true,
				Strict: true,
			},
		)
		err := serializer.Encode(v.Object, writer)
		if err != nil {
			fmt.Println(err)
		}
		writer.Flush()

		var result lib.Converted
		result.Name = name
		result.Path = "/"
		result.Payload = *buff

		converted = append(converted, result)
	}

	return converted
}
