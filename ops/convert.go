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

package ops

import (
	"log"
	generators "shifter/generators"
	inputs "shifter/inputs"
)


/*
	TODO
	- Add Errors Handling to Convert,
	- Catch Convert Errors,
	- Return error struct on Errors
*/
func Convert(inputType string, filename string, generator string, output string, flags map[string]string) {

	switch inputType {
	case "template":
		t := inputs.Template(filename)
		switch generator {
		case "helm":
			generators.Helm(output, t)
		}
	case "yaml":
		t := inputs.Yaml(filename, flags)
		switch generator {
		case "yaml":
			generators.Yaml(output, t)
		}
	case "cluster":
		log.Fatal("Openshift resources have not been implemented yet!")
	}
	return
}
