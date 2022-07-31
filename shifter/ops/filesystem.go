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
	log.Println("-------------------------------------------------------------")
	log.Println("Storage Type:		", fileObj.StorageType)
	log.Println("Source Path:		", fileObj.Path)
	log.Println("File Extention:		", fileObj.Ext)
	log.Println("-------------------------------------------------------------")
}

// Returns FileObject Content (bytes.Buffer) as a String
func (fileObj *FileObject) ContentAsString() string {
	return fileObj.Content.String()
}

// Writes the Contents of a FileObject to the Correct Storage Type
func (fileObj *FileObject) WriteFile() {
	// Path is GCS Bucket
	if fileObj.StorageType == GCS {
		// Traverse and Create Files Objects for GCS Bucket
		fileObj.WriteGCSFile()
	} else {
		// Traverse and Create Files Objects for Local Directory
		fileObj.WriteLCLFile()
	}
}

// Loads File from Storage based on Storage Type and Stores Files content as bytes.Buffer
func (fileObj *FileObject) LoadFile() {
	// Path is GCS Bucket
	if fileObj.StorageType == GCS {
		// Traverse and Create Files Objects for GCS Bucket
		fileObj.LoadGCSFile()
	} else {
		// Traverse and Create Files Objects for Local Directory
		fileObj.LoadLCLFile()
	}
}

// Takes Source Path, Determines Storage Tyoe and Creates File List
func ProcessPath(path string) ([]*FileObject, error) {
	// Path is GCS Bucket
	if strings.Contains(path, "gs://") {
		// Traverse and Create Files Objects for GCS Bucket
		return ProcessGCSPath(path)
	} else {
		// Traverse and Create Files Objects for Local Directory
		return ProcessLCLPath(path)
	}
}
