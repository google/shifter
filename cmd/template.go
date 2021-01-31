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
)

type Template struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		CreationTimestamp string `yaml:"creationTimestamp"`
		Name              string `yaml:"name"`
	}
	Objects    []map[interface{}]interface{}
	Parameters []struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		Required    bool   `yaml:"required"`
		Value       string `yaml:"value"`
	}
}

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Convert openshift templates to helm charts",
	Long: `Convert an openshift template to a helm chart

Usage: shifter template -i ./input.yaml -o ./output_dir
Supply the input file with the -i or --input flag
Supply the output using the -o or --output flag, the directory will be created with the contents of the helm chart.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("templates called")
		readYaml(args[0])
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)
	templateCmd.Flags().BoolP("input", "i", false, "Path to the input file to covert, must be in Openshift format")
	templateCmd.Flags().BoolP("output", "o", false, "Path to the output file for the results on the conversion")
}

func readYaml(input string) {
	yamlFile, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Println(err)
	}
	t := Template{}
	err = yaml.Unmarshal(yamlFile, &t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
	var o []map[interface{}]interface{}
	o = t.Objects
	fmt.Println(o)
}
