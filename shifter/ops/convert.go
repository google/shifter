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
	"path/filepath"
	"shifter/generator"
	"shifter/input"
	"shifter/lib"
	"strings"

	"github.com/google/uuid"
)

type LogObject struct {
}

type Converter struct {
	UUID       string // Unique ID of the Run
	InputType  string
	SourcePath string
	Generator  string
	OutputPath string
	Flags      map[string]string

	SourceFiles []*FileObject
	OutputFiles []*FileObject

	Logs []*LogObject
}

type DownloadFile struct {
	Link     string `json:"link"`
	Filename string `json:"filename"`
}

// Input Types
const YAML string = "YAML"
const TEMPLATE string = "TEMPLATE"

// Create New Converter
func NewConverter(inputType string, sourcePath string, generator string, outputPath string, flags map[string]string) (*Converter, error) {
	log.Printf("🧰 💡 INFO: Creating New Shifter Converter")
	// Create New Instance of Converter
	converter := &Converter{}

	// Create UUID for Converter
	converter.UUID = uuid.New().String()

	// Set all the Variables for the Converter
	converter.InputType = inputType
	converter.SourcePath = sourcePath
	converter.Generator = generator
	converter.OutputPath = outputPath
	converter.Flags = flags

	// Process the Path and Create Array of File Objects
	files, err := ProcessPath(converter.SourcePath)
	if err != nil {
		// Error: Processing File Source Path
		log.Printf("🧰 ❌ ERROR: Processing File Source Path.")
		return converter, err
	} else {
		// Success: Processing File Source Path
		log.Printf("🧰 ✅ SUCCESS: Processing File Source Path")
	}

	// Set Converter Files
	converter.SourceFiles = files
	if len(converter.SourceFiles) > 0 {
		err := converter.LoadSourceFiles()
		if err != nil {
			// Error: Error Loading Source Files
			return converter, err
		}
	}

	// Success Creating Shifter Converter
	return converter, nil
}

func (converter *Converter) WriteSourceFiles() error {
	log.Printf("🧰 💡 INFO: Writing all Source Files")
	// Process Input Objects
	for _, file := range converter.SourceFiles {
		err := file.WriteFile()
		if err != nil {
			// Error: Error Writing File
			return err
		}
	}
	// Success - Writing Source Files
	return nil
}

func (converter *Converter) LoadSourceFiles() error {
	log.Printf("🧰 💡 INFO: Loading all Source Files")
	// Process Input Objects
	for _, file := range converter.SourceFiles {
		err := file.LoadFile()
		if err != nil {
			// Error: Error Loading File
			return err
		}
	}
	// Success - Loading Source Files
	return nil
}

func (converter *Converter) ListSourceFiles() error {
	log.Printf("🧰 💡 INFO: Listing all Source Files")
	// Process Input Objects
	for _, file := range converter.SourceFiles {
		file.Meta()
	}
	// Success - Listing Source Files
	return nil
}

func (converter *Converter) ListOutputFiles() error {
	log.Printf("🧰 💡 INFO: Listing all Output Files")
	// Process Input Objects
	for _, file := range converter.OutputFiles {
		file.Meta()
	}
	// Success - Listing Output Files
	return nil
}

func (converter *Converter) ConvertFiles() error {
	// Process Input Objects
	for _, file := range converter.SourceFiles {

		var (
			resources  []lib.Converted
			err        error
			sourceFile []lib.K8sobject
			values     []lib.OSTemplateParams
		)

		switch strings.ToUpper(converter.InputType) {
		// Input Type == YAML
		case YAML:
			sourceFile, err := input.Yaml(file.Content, converter.Flags)
			if err != nil {
				// Error: Parsing Input YAML
				log.Printf("🧰 ❌ ERROR: Parsing Input YAML")
				return err
			}
			resources, err = generator.NewGenerator(converter.Generator, file.Filename, sourceFile)
			if err != nil {
				// Error: Unable to Create Shifter 'YAML' Generator
				log.Printf("🧰 ❌ ERROR: Create 'YAML' Shifter Generator.")
				return err
			} else {
				// Succes: Creating Shifter 'YAML' Generator
				log.Printf("🧰 ✅ SUCCESS: Shifter 'YAML' Generator Successufly Created.")
			}
		// Input Type == TEMPLATE
		case TEMPLATE:
			sourceFile, values, err = input.Template(file.Content, converter.Flags)
			if err != nil {
				// Error: Parsing Input TEMPLATE
				log.Printf("🧰 ❌ ERROR: Parsing Input TEMPLATE")
				return err
			}
			resources, err = generator.NewGenerator(converter.Generator, file.Filename, sourceFile, values)
			if err != nil {
				// Error: Unable to Create Shifter 'Template' Generator
				log.Printf("🧰 ❌ ERROR: Create 'Template' Shifter Generator.")
				return err
			} else {
				// Succes: Creating Shifter 'Template' Generator
				log.Printf("🧰 ✅ SUCCESS: Shifter 'Template' Generator Successufly Created.")
			}
		}

		//outputFileName := fmt.Sprint(idx)
		for k := range resources {
			log.Printf("🧰 💡 INFO: Creating Shifter File Object")
			fileObj := &FileObject{
				StorageType: file.StorageType,
				//SourcePath:    (converter.OutputPath + "/" + r[k].Path + r[k].Name + filepath.Ext(file.SourcePath)),
				Path:          (converter.OutputPath + "/" + resources[k].Path + resources[k].Name),
				Filename:      file.Filename,
				Ext:           filepath.Ext(file.Path),
				Content:       resources[k].Payload,
				ContentLength: file.ContentLength,
			}

			// Write Converted File to Storage
			err := fileObj.WriteFile()
			if err != nil {
				// Error: Error Writing File
				return err
			}

			// Add Converted File Object to Converter
			converter.OutputFiles = append(converter.OutputFiles, fileObj)
		}
	}

	// Success Converting Files
	return nil
}

/*
TODO - Remove this Function
func (converter *Converter) BuildDownloadFiles() ([]*DownloadFile, error) {
	var files []*DownloadFile

	// Process Output Objects
	for _, file := range converter.OutputFiles {
		dlFile := &DownloadFile{}
		dlFile.Link = "https://somefile.com"
		dlFile.Filename = file.Filename
		files = append(files, dlFile)
	}

	// Success Building Download Files
	return files, nil
}*/
