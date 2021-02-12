/*
Copyright 2019 Google LLC
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"shifter/generator"
)

type Template struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		CreationTimestamp string                 `yaml:"creationTimestamp"`
		Name              string                 `yaml:"name"`
		Annotations       map[string]interface{} //annotations have a unkown structure so we use a generic interface
	}
	Objects []struct {
		ApiVersion string                 `yaml:"apiVersion"`
		Kind       string                 `yaml:"kind"`
		Metadata   map[string]interface{} //metadata has a unkown structure so we use a generic interface
		Spec       map[string]interface{} //specs are dependent on the kind so we use a generic interface
		Data       map[string]interface{} //specs are dependent on the kind so we use a generic interface
	}
	Parameters []struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		Required    bool   `yaml:"required"`
		Value       string `yaml:"value"`
	}
}

var (
	input  string
	output string
	kind   string
)

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Convert openshift templates to helm charts",
	Long: `Convert an openshift template to a helm chart

Usage: shifter template -i ./input.yaml -o ./output_dir
Supply the input file with the -i or --input flag
Supply the output using the -o or --output flag, the directory will be created with the contents of the helm chart.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Shifter - Templates")


		t := readYaml(input)
		parseOS(t)
		//var o []map[interface{}]interface{}
		//o = t.Objects
		//fmt.Println(o)
		//fmt.Println(t.Parameters)
		//for k, v := range o {
		//	fmt.Println(k, v)
		//}
		/*fmt.Println(o[1])
		for k, v := range o[1] {
			fmt.Println(k, v)
		}
		*/
		switch kind {
		case "helm":
			generator.CreateChart(output)
		}
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)
	templateCmd.Flags().StringVarP(&input, "input", "i", "", "Path to the input file to covert, must be in Openshift format")
	templateCmd.Flags().StringVarP(&kind, "kind", "k", "helm", "Output kind options are either helm or kpt")
	templateCmd.Flags().StringVarP(&output, "output", "o", "", "Path to the output file for the results on the conversion")
}

func readYaml(file string) Template {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	t := Template{}
	err = yaml.Unmarshal(yamlFile, &t)
	if err != nil {
		fmt.Println(err)
	}
	return t
}

func parseOS(t Template) {
	fmt.Println("*******************************************")
	fmt.Println(t)
	fmt.Println("*******************************************")

	for i, o := range t.Objects {
		switch o.Kind {
		case "DeploymentTemplate":
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		case "ImageStream":
			fmt.Println("XXXXXXXXXXXXXXXXXXX")
		}
		fmt.Println(i, o.Kind)
		for g, h := range o.Spec {
			fmt.Println(g, h)
		}
	}
}
