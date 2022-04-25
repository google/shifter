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
	"os"
	"path/filepath"
)

func (fileObj *FileObject) WriteLCLFile() {

	if _, err := os.Stat(fileObj.SourcePath); os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(fileObj.SourcePath), 0700) // Create your file
	}

	// Create New File
	f, err := os.Create(fileObj.SourcePath)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	// Write the Bytes Buffer to File
	_, err = f.Write(fileObj.Content.Bytes())
	if err != nil {
		log.Println(err)
		return
	}
	f.Sync()
}

func (fileObj *FileObject) LoadLCLFile() {

	file, err := os.Open(fileObj.SourcePath)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		log.Println(err)
		return
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	bytesread, err := file.Read(buffer)
	if err != nil {
		log.Println(err)
		return
	}

	// Set FileObject Content & Length MetaData
	fileObj.Content = *bytes.NewBuffer(buffer)
	fileObj.ContentLength = bytesread
}

func ProcessLCLPath(path string) ([]*FileObject, error) {
	var files []*FileObject
	//Get File Info on Path
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	// Check File or Directory
	switch mode := fileInfo.Mode(); {

	// Is Directory
	case mode.IsDir():
		err := filepath.Walk(path, func(filePath string, f os.FileInfo, err error) error {
			if !f.IsDir() {
				// Create File Object for every file in Directory
				fileObj := &FileObject{
					StorageType: LCL,
					SourcePath:  filePath,
					Ext:         filepath.Ext(filePath),
				}
				// Add File Object to Array of Files
				files = append(files, fileObj)
			}
			return err
		})
		if err != nil {
			log.Println(err)
			return files, err
		}
		return files, err

	// Is A File
	case mode.IsRegular():
		// Create File Object for file
		fileObj := &FileObject{
			StorageType: LCL,
			SourcePath:  path,
			Ext:         filepath.Ext(path),
		}
		// Add File Object to Array of Files
		files = append(files, fileObj)
	}

	// Return Array of Files
	return files, nil
}
