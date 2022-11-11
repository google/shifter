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

// Create New Converter
func NewConverter(inputType string, sourcePath string, generator string, outputPath string, flags map[string]string) (*Converter, error) {
	converter := &Converter{}
	converter.UUID = uuid.New().String()

	converter.InputType = inputType
	converter.SourcePath = sourcePath
	converter.Generator = generator
	converter.OutputPath = outputPath
	converter.Flags = flags

	files, err := ProcessPath(converter.SourcePath)
	if err != nil {
		lib.CLog("error", "Processing file source path", err)
		return converter, err
	}

	converter.SourceFiles = files
	if len(converter.SourceFiles) > 0 {
		err := converter.LoadSourceFiles()
		if err != nil {
			lib.CLog("error", "Loading source files", err)
			return converter, err
		}
	}

	return converter, nil
}

func (converter *Converter) WriteSourceFiles() error {
	lib.CLog("info", "Writing to destination files")
	for _, file := range converter.SourceFiles {
		err := file.WriteFile()
		if err != nil {
			lib.CLog("error", "Loading source files", err)
			return err
		}
	}
	return nil
}

func (converter *Converter) LoadSourceFiles() error {
	lib.CLog("info", "Reading all source files")
	// Process Input Objects
	for _, file := range converter.SourceFiles {
		err := file.LoadFile()
		if err != nil {
			lib.CLog("error", "Loading source files", err)
			return err
		}
	}
	return nil
}

func (converter *Converter) ListSourceFiles() error {
	lib.CLog("info", "Listing all source files")
	// Process Input Objects
	for _, file := range converter.SourceFiles {
		file.Meta()
	}
	return nil
}

func (converter *Converter) ListOutputFiles() error {
	log.Printf("🧰 💡 INFO: Listing all Output Files")
	// Process Input Objects
	for _, file := range converter.OutputFiles {
		file.Meta()
	}
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

		switch strings.ToLower(converter.InputType) {
		// Input Type == YAML
		case "yaml":
			sourceFile, err := input.Yaml(file.Content, converter.Flags)
			if err != nil {
				lib.CLog("error", "Parsing yaml", err)
				return err
			}
			resources, err = generator.NewGenerator(converter.Generator, file.Filename, sourceFile)
			if err != nil {
				lib.CLog("error", "Unable to create generator "+converter.Generator, err)
				return err
			}
		case "template":
			sourceFile, values, err = input.Template(file.Content, converter.Flags)
			if err != nil {
				lib.CLog("error", "Parsing template", err)
				return err
			}
			resources, err = generator.NewGenerator(converter.Generator, file.Filename, sourceFile, values)
			if err != nil {
				lib.CLog("error", "unable to create generator "+converter.Generator, err)
				return err
			}
		}

		for k := range resources {
			fileObj := &FileObject{
				StorageType: file.StorageType,
				//SourcePath:    (converter.OutputPath + "/" + r[k].Path + r[k].Name + filepath.Ext(file.SourcePath)),
				Path:          (converter.OutputPath + "/" + resources[k].Path + resources[k].Name),
				Filename:      file.Filename,
				Ext:           filepath.Ext(file.Path),
				Content:       resources[k].Payload,
				ContentLength: file.ContentLength,
			}

			err := fileObj.WriteFile()
			if err != nil {
				// Error: Error Writing File
				lib.CLog("error", "Unable to write to file", err)
				return err
			}

			converter.OutputFiles = append(converter.OutputFiles, fileObj)
		}
	}

	return nil
}
