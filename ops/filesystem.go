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
	"fmt"
	"log"
	"os"
	"path/filepath"
	"shifter/lib"
	"bytes"
)

type JSONResponse struct {
	Value1 string `json:"Filename"`
	Value2 string `json:"Link"`
}

func CreateDir(srcPath string) {
	if _, err := os.Stat(srcPath); os.IsNotExist(err) {
		os.MkdirAll(srcPath, 0700)
	}
}

func GetFilename(path string) string {
	return filepath.Base(path)
}
func GetFileLink(uuid string, path string) string {
	return ("/download/" + uuid + "/" + filepath.Base(path))
}

func GetFiles(uuid string, srcPath string) []File {

	// Create Splice of File Structs
	var files []File

	// Walk Upload Directory Directory
	err := filepath.Walk(srcPath, func(path string, info os.FileInfo, err error) error {
		// Add Each File to Splice
		files = append(files, File{Filename: GetFilename(path), Link: GetFileLink(uuid, path)})
		return nil
	})
	if err != nil {
		panic(err)
	}

	// Return Splice/Array of Files
	return files
}


/*
	Cross Platform File Object. 
	Can be used to Store Content and MetaData about
	Files stored in Local Storage and or GCS Storage where
	contents is written as a pointer to a bytes buffer
*/
type FileObject struct {
	UUID 			string			// Unique ID of the Run
	StorageType 	string			// GCS or LCL Storage
	Root			string			// Root Directory "/data" LCL Storage
	Bucket			string			// Root Bucket Nmae GCS Storage
	Path   			string			// Sub Path ["raw", "output"]
	Filename		string			// Filename
	Ext				string			// File Extention
	Content			bytes.Buffer	// Content as Bytes Buffer
	ContentLength	int				// Content Length (len(bytes.buffer))
}


func WriteFile(fileObj FileObject) error{

	// Handle Writing File to GCS
	if fileObj.StorageType == "GCS"{
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
}