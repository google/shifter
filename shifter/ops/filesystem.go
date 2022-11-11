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
	"bytes"
	"log"
	"strings"
)

type FileObject struct {
	UUID          string       // Unique ID of the Run
	StorageType   string       // GCS or LCL Storage
	Path          string       // Bucket or Local Path
	Ext           string       // File Extention
	Filename      string       // Filename
	Content       bytes.Buffer // Content as Bytes Buffer
	ContentLength int          // Content Length (len(bytes.buffer))
}

/*
Cross Platform File Object.
Can be used to Store Content and MetaData about
Files stored in Local Storage and or GCS Storage where
contents is written as a pointer to a bytes buffer
*/
const GCS string = "GCS"
const LCL string = "LCL"

// Display File Object Content for Debugging.
func (fileObj *FileObject) Meta() {
	// Log File Object Data.
	log.Printf("🌐 📜 INFO: Shifter File Object \n[\nStorage Type: %s, \nFilename: %s, \nExtension: %s, \nPath: %s \n]",
		fileObj.StorageType,
		fileObj.Filename,
		fileObj.Ext,
		fileObj.Path,
	)
}

// Returns FileObject Content (bytes.Buffer) as a String
func (fileObj *FileObject) ContentAsString() string {
	return fileObj.Content.String()
}

// Writes the Contents of a FileObject to the Correct Storage Type
func (fileObj *FileObject) WriteFile() error {
	// Path is GCS Bucket
	if fileObj.StorageType == GCS {
		// Traverse and Create Files Objects for GCS Bucket
		err := fileObj.WriteGCSFile()
		if err != nil {
			// Error Loading File. Bubble Error Up
			return err
		}
	} else {
		// Traverse and Create Files Objects for Local Directory
		err := fileObj.WriteLCLFile()
		if err != nil {
			// Error Loading File. Bubble Error Up
			return err
		}
	}
	// Success Writing File
	return nil
}

// Loads File from Storage based on Storage Type and Stores Files content as bytes.Buffer
func (fileObj *FileObject) LoadFile() error {
	// Path is GCS Bucket
	if fileObj.StorageType == GCS {
		// Traverse and Load Files Objects for GCS Bucket
		err := fileObj.LoadGCSFile()
		if err != nil {
			// Error Loading File. Bubble Error Up
			return err
		}
	} else {
		// Traverse and Load Files Objects for Local Directory
		err := fileObj.LoadLCLFile()
		if err != nil {
			// Error Loading File. Bubble Error Up
			return err
		}
	}
	// Success Loading File
	return nil
}

// Takes Source Path, Determines Storage Tyoe and Creates File List
func ProcessPath(path string) ([]*FileObject, error) {
	var files []*FileObject

	// Path is GCS Bucket
	if strings.Contains(path, "gs://") {
		// Traverse and Create Files Objects for GCS Bucket
		fileObj, err := ProcessGCSPath(path)
		if err != nil {
			// Error Loading GCS File. Bubble Error Up
			return files, err
		} else {
			// Success Loading GCS File.
			return fileObj, err
		}
	} else {
		// Traverse and Create Files Objects for Local Directory
		fileObj, err := ProcessLCLPath(path)
		if err != nil {
			// Error Loading Local File. Bubble Error Up
			return files, err
		} else {
			// Success Loading Local File.
			return fileObj, err
		}
	}
}
