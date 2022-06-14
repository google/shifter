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
	"shifter/lib"
)

type Generator struct {
	OutputType string
	Name       string
	Input      struct {
		Object     []lib.K8sobject
		Parameters []lib.OSTemplateParams
	}
}

type Input struct {
	Object []lib.K8sobject
}

func NewGenerator(outputType string, name string, input []lib.K8sobject, parameters []lib.OSTemplateParams) []lib.Converted {
	generator := &Generator{}

	outputType = outputType
	generator.Name = name
	generator.Input.Object = input
	generator.Input.Parameters = parameters

	switch outputType {
	case "yaml":
		return generator.Yaml(name, generator.Input.Object)
	case "helm":
		return generator.Helm(name, generator.Input.Object, generator.Input.Parameters)
	}

	return nil
}
