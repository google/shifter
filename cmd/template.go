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

type kube struct {
	Parameters []struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		Required    bool   `yaml:"required"`
		Value       string `yaml:"value"`
	}
	Objects []struct {
		ApiVersion string                 `yaml:"apiVersion"`
		Kind       string                 `yaml:"kind"`
		Metadata   map[string]interface{} //metadata has a unkown structure so we use a generic interface
		Spec       map[string]interface{} //specs are dependent on the kind so we use a generic interface
		Data       map[string]interface{} //specs are dependent on the kind so we use a generic interface
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

		k, err := yaml.Marshal(parseOS(t))
		if err != nil {
			fmt.Println(err)
		}

		switch kind {
		case "helm":
			generator.Generate(output, k)
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
	template := Template{}
	err = yaml.Unmarshal(yamlFile, &template)
	if err != nil {
		fmt.Println(err)
	}
	return template
}

func parseOS(t Template) kube {
	var k8s kube

	//iterate over the objects and modify them as needed
	for _, o := range t.Objects {
		switch o.Kind {
		case "DeploymentConfig":
			o.Kind = "Deployment"
			o.ApiVersion = "apps/v1"
			k8s.Objects = append(k8s.Objects, o)
		case "ImageStream":
		case "Route":
		default:
			k8s.Objects = append(k8s.Objects, o)
		}
	}

	for _, y := range t.Parameters {
		k8s.Parameters = append(k8s.Parameters, y)
	}
	return k8s
}
