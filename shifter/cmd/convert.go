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
	//filename   string
	sourcePath string
	outputPath string
	generator  string
	pFlags     []string
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
Supply the input file or directory of files with the -i or --input-format flag
Supply the output using the -o or --output-path flag, the directory will be created with the contents of the helm chart.
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
		log.Println("Converting", inputType, sourcePath, "to", generator, outputPath)
		flags := ProcFlags(pFlags)
		con := ops.NewConverter(inputType, sourcePath, generator, outputPath, flags)
		con.ConvertFiles()
		//ops.Convert(inputType, filename, generator, output, flags)
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().StringVarP(&inputType, "input-format", "i", "yaml", "Input format. One of: yaml|template")
	convertCmd.Flags().StringVarP(&generator, "output-format", "t", "", "Output format. One of: yaml|helm")
	convertCmd.Flags().StringVarP(&sourcePath, "source-path", "f", "", "Relative Local Path (./data/source) or Google Cloud Storage Bucket Path (gs://XXXXXXX/source/) to convert (contents must be in OpenShift format)")
	convertCmd.Flags().StringVarP(&outputPath, "output-path", "o", "", "Relative Local Path (./data/output) or Google Cloud Storage Bucket Path (gs://XXXXXXX/output/) for Converted Files to be Written")
	convertCmd.Flags().StringSliceVarP(&pFlags, "pflags", "", []string{}, "Flags passed to the processor")
	convertCmd.MarkFlagRequired("source-path")
	convertCmd.MarkFlagRequired("output-path")
	//convertCmd.Flags().StringVarP(&filename, "filename", "f", "", "Path to the file or directory to convert (contents must be in OpenShift format)")
	//convertCmd.Flags().StringVarP(&output, "output-path", "o", "", "Relative path to the output directory for the results on the conversion")
	//convertCmd.MarkFlagRequired("filename")
}
