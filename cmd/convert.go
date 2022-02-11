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
	"log"
	ops "shifter/ops"

	"github.com/spf13/cobra"
)

var (
	inputType string
	filename  string
	output    string
	generator string
	pFlags    []string
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert Openshift Resources to Kubernetes native formats",
	Long: `

   _____ __    _ ______           
  / ___// /_  (_) __/ /____  _____
  \__ \/ __ \/ / /_/ __/ _ \/ ___/
 ___/ / / / / / __/ /_/  __/ /    
/____/_/ /_/_/_/  \__/\___/_/     
                                  

Convert OpenShift resources to kubernetes native formats

Usage: shifter convert -i ./input.yaml -o ./output_dir -k kind -t kind
Supply the input file or directory of files with the -i or --input flag
Supply the output using the -o or --output flag, the directory will be created with the contents of the helm chart.
`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(`
   _____ __    _ ______           
  / ___// /_  (_) __/ /____  _____
  \__ \/ __ \/ / /_/ __/ _ \/ ___/
 ___/ / / / / / __/ /_/  __/ /    
/____/_/ /_/_/_/  \__/\___/_/     
                                 
----------------------------------------
			`)
		log.Println("Converting", inputType, filename, "to", generator, output)

		flags := procFlags(pFlags)
		log.Println("Processor Flags:", flags)
		switch inputType {
		case "template":
			t, p, n := inputs.Template(filename, flags)
			switch generator {
			case "helm":
				generators.Helm(output, t, p, n)
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
		log.Println("Conversion completed.")
		flags := ProcFlags(pFlags)
		//"yaml ./_test/yaml/multidoc/os-nginx.yaml yaml ./output map[]"
		ops.Convert(inputType, filename, generator, output, flags)
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().StringVarP(&inputType, "input-format", "i", "yaml", "Input format. One of: yaml|template")
	convertCmd.Flags().StringVarP(&filename, "filename", "f", "", "Path to the file or directory to convert (contents must be in OpenShift format)")
	convertCmd.Flags().StringVarP(&generator, "output-format", "t", "", "Output format. One of: yaml|helm")
	convertCmd.Flags().StringVarP(&output, "output-path", "o", "", "Relative path to the output directory for the results on the conversion")
	convertCmd.Flags().StringSliceVarP(&pFlags, "pflags", "", []string{}, "Flags passed to the processor")
	convertCmd.MarkFlagRequired("filename")
	convertCmd.MarkFlagRequired("output-path")
}
