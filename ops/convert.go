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
	"fmt"
	"log"
	generators "shifter/generators"
	inputs "shifter/inputs"
	lib "shifter/lib"
)

/*
	TODO
	- Add Errors Handling to Convert,
	- Catch Convert Errors,
	- Return error struct on Errors
*/
func Convert(inputType string, filename string, generator string, output string, flags map[string]string) {
	var outputFiles []lib.Converted
	switch inputType {
	case "template":
		name, t, p := inputs.Template(filename, flags)
		switch generator {
		case "helm":
			outputFiles = generators.NewGenerator(generator, name, t, p)
		}
	case "yaml":
		name, t := inputs.Yaml(filename, flags)
		switch generator {
		case "yaml":
			outputFiles = generators.NewGenerator(generator, name, t, nil)
		}
	case "cluster":
		log.Fatal("Openshift resources have not been implemented yet!")
	}

	for k := range outputFiles {
		fmt.Println(outputFiles[k].Path + outputFiles[k].Name)
		fmt.Println(outputFiles[k].Payload.String())
	}
	return
}
