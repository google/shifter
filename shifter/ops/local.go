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

// Write bytes.Buffer to Local File System
func (fileObj *FileObject) WriteLCLFile() error {
	log.Printf("🧰 📜 INFO: Writing Shifter File Object to Local Filesystem")

	// Check if Output Directory Exits, IF Not, Make output Directory
	if _, err := os.Stat(fileObj.Path); os.IsNotExist(err) {
		log.Printf("🧰 📁 INFO: Proposed Output Directory Does Not Exist, Creating Directory")
		err := os.MkdirAll(filepath.Dir(fileObj.Path), 0700) // Create output directory
		if err != nil {
			// Error: Unable to Create Missing Output Directory
			log.Printf("🧰 ❌ ERROR: Unable to Create Output Directory: '%s'.", fileObj.Path)
			return err
		} else {
			// Success: Created Missing Output Directory
			log.Printf("🧰 ✅ SUCCESS: Created Missing Output Directory: '%s'.", fileObj.Path)
		}
	}

	// Create New File Name
	fileName := fileObj.Path + "." + fileObj.Ext

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0700)
	if err != nil {
		// Error: Unable to Create New File
		log.Printf("🧰 ❌ ERROR: Creating New File: '%s'.", fileName)
		return err
	} else {
		// Success: Creating New File
		log.Printf("🧰 ✅ SUCCESS: Creating New File: '%s'.", fileName)
	}
	defer f.Close()
	f.WriteString("---\n")
	// Write the Bytes Buffer to File
	_, err = f.Write(fileObj.Content.Bytes())
	if err != nil {
		// Error: Unable to Write New File
		log.Printf("🧰 ❌ ERROR: Writing Content to File: '%s'.", fileName)
		return err
	} else {
		// Success: Writing New File
		log.Printf("🧰 ✅ SUCCESS: Writing Content to File: '%s'.", fileName)
	}
	f.Sync()

	// Successfull Writen file to Local File System
	log.Printf("🧰 ✅ SUCCESS: File written to Local File System: '%s'.", fileName)
	return nil
}

func (fileObj *FileObject) LoadLCLFile() error {
	log.Printf("🧰 💡 INFO: Loading Shifter File Object from Local Filesystem")

	file, err := os.Open(fileObj.Path)
	if err != nil {
		log.Printf("🧰 ❌ ERROR: Opening file on Local File System: '%s'.", fileObj.Path)
		return err
	} else {
		log.Printf("🧰 ✅ SUCCESS: Opening file on Local File System: '%s'.", fileObj.Path)
	}

	log.Printf("🧰 💡 INFO: Reading Shifter File - '%s'", fileObj.Path)
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		// ERROR: Obtaining Statistical Information about Local File System File
		log.Printf("🧰 ❌ ERROR: Getting file information: '%s'.", fileObj.Path)
		return err
	} else {
		// SCUCCESS: Obtaining Statistical Information about Local File System File
		log.Printf("🧰 ✅ SUCCESS: Obtaining file information: '%s'.", fileObj.Path)
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	bytesread, err := file.Read(buffer)
	if err != nil {
		// ERROR: Reading file contents into Bytes Buffer
		log.Printf("🧰 ❌ ERROR: Reading file contents into Buffer.")
		return err
	} else {
		// Success: Reading file contents into Buffer
		log.Printf("🧰 ✅ SUCCESS: Reading file contents into Buffer.")
	}

	// Set FileObject Content & Length MetaData
	fileObj.Content = *bytes.NewBuffer(buffer)
	fileObj.ContentLength = bytesread

	// SCUCCESS: Reading Local File System File
	log.Printf("🧰 ✅ SUCCESS: Reading Shifter File Object from Local Filesystem.")
	return nil
}

func ProcessLCLPath(path string) ([]*FileObject, error) {
	var files []*FileObject
	//Get File Info on Path
	fileInfo, err := os.Stat(path)
	if err != nil {
		// ERROR: Obtaining Statistical Information about Local File System File
		log.Printf("🧰 ❌ ERROR: Getting file information: '%s'.", path)
		return files, err
	} else {
		// SCUCCESS: Obtaining Statistical Information about Local File System File
		log.Printf("🧰 ✅ SUCCESS: Obtaining file information: '%s'.", path)
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
					Path:        filePath,
					Ext:         filepath.Ext(filePath),
					Filename:    filepath.Base(filePath),
				}
				// Add File Object to Array of Files
				files = append(files, fileObj)
			}
			// Bubble Up File Walk Error
			return err
		})
		if err != nil {
			// Error: Traversing Filepath
			log.Printf("🧰 ❌ ERROR: Unable to Traverse Filepath.")
			return files, err
		}
		//Success Processing Directory in FileWalk.
		return files, err

	// Is A File
	case mode.IsRegular():
		// Create File Object for file
		fileObj := &FileObject{
			StorageType: LCL,
			Path:        path,
			Ext:         filepath.Ext(path),
			Filename:    filepath.Base(path),
		}
		// Add File Object to Array of Files
		files = append(files, fileObj)
	}

	// Success Processing LCL File Path - Return Array of Files
	return files, nil
}
