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
	"strings"
)

type JSONResponse struct {
	Value1 string `json:"Filename"`
	Value2 string `json:"Link"`
}

//func CreateDir(srcPath string) {
//	if _, err := os.Stat(srcPath); os.IsNotExist(err) {
//		os.MkdirAll(srcPath, 0700)
//	}
//}

//func GetFilename(path string) string {
//	return filepath.Base(path)
//}

//func GetFileLink(uuid string, path string) string {
//	return ("/download/" + uuid + "/" + filepath.Base(path))
//}

//func GetFiles(uuid string, srcPath string) []File {

// Create Splice of File Structs
//	var files []File

// Walk Upload Directory Directory
//	err := filepath.Walk(srcPath, func(path string, info os.FileInfo, err error) error {
// Add Each File to Splice
//		files = append(files, File{Filename: GetFilename(path), Link: GetFileLink(uuid, path)})
//		return nil
//	})
//	if err != nil {
//		panic(err)
//	}

// Return Splice/Array of Files
//	return files
//}

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
	log.Println("Source Path:		", fileObj.SourcePath)
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

/*
func WriteFile(fileObj FileObject) error {
	// Handle Writing File to GCS
	if fileObj.StorageType == "GCS" {
		log.Println("Writting File to GCS Bucket")
		// Run GCSL Stream File Upload Function
		err := lib.GCSStreamFileUpload(fileObj.Content, fileObj.Bucket, fmt.Sprintf("%s/%s/%s", fileObj.Path, fileObj.UUID, fileObj.Filename))
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		log.Println("Writting File to Local Storage")
		err := lib.LocalStreamFileUpload(fileObj.Content, fmt.Sprintf("data/%s/%s/", fileObj.Path, fileObj.UUID), fileObj.Filename)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
	//log.Println(fileObj)
}*/
