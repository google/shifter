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
	"os"
	"path/filepath"
	"shifter/lib"
)

// Write bytes.Buffer to Local File System
func (fileObj *FileObject) WriteLCLFile() error {
	lib.CLog("debug", "Writing file object to local file system")

	// Check if Output Directory Exits, IF Not, Make output Directory
	lib.CLog("debug", "Path: "+fileObj.Path)

	if _, err := os.Stat(fileObj.Path); os.IsNotExist(err) {
		lib.CLog("info", "Output directory does not exist... creating")
		err := os.MkdirAll(fileObj.Path, 0700)
		if err != nil {
			lib.CLog("error", "Unable to create output directory", err)
			return err
		}
	}

	var fileName string

	if filepath.Ext(fileObj.Filename) == "" {
		fileName = filepath.Join(fileObj.Filename, fileObj.Ext)
	} else {
		fileName = fileObj.Filename
	}

	fileName = filepath.Join(fileObj.Path, fileName)

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		lib.CLog("error", "creating new file"+fileName, err)
		return err
	}

	defer f.Close() // TODO - Address goSec Error

	f.WriteString("---\n")
	_, err = f.Write(fileObj.Content.Bytes())
	if err != nil {
		lib.CLog("error", "Writing content to file "+fileName, err)
		return err
	}

	err = f.Sync()
	if err != nil {
		lib.CLog("error", "Writing content to file "+fileName, err)
		return err
	}

	lib.CLog("debug", "Files written to local file system")
	return nil
}

func (fileObj *FileObject) LoadLCLFile() error {
	file, err := os.Open(fileObj.Path)
	if err != nil {
		lib.CLog("error", "Opening file from local file system: "+fileObj.Path, err)
		return err
	}

	lib.CLog("debug", "Reading file - "+fileObj.Path)
	defer file.Close() // TODO - Address goSec Error

	fileinfo, err := file.Stat()
	if err != nil {
		lib.CLog("error", "Getting file information for file: "+fileObj.Path, err)
		return err
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	bytesread, err := file.Read(buffer)
	if err != nil {
		lib.CLog("error", "Error reading source file", err)
		return err
	} else {
		lib.CLog("debug", "File contents loaded into Buffer.")
	}

	fileObj.Content = *bytes.NewBuffer(buffer)
	fileObj.ContentLength = bytesread

	return nil
}

func ProcessLCLPath(path string) ([]*FileObject, error) {
	var files []*FileObject
	fileInfo, err := os.Stat(path)
	if err != nil {
		lib.CLog("error", "Unable to get file information for "+path, err)
		return files, err
	}
	lib.CLog("debug", "Reading file information from: "+path)

	switch mode := fileInfo.Mode(); {

	case mode.IsDir():
		err := filepath.Walk(path, func(filePath string, f os.FileInfo, err error) error {
			if !f.IsDir() {
				fileObj := &FileObject{
					StorageType: LCL,
					Path:        filePath,
					Ext:         filepath.Ext(filePath),
					Filename:    filepath.Base(filePath),
				}
				files = append(files, fileObj)
			}
			return err
		})
		if err != nil {
			lib.CLog("error", "Unable to traverse filepath", err)
			return files, err
		}
		return files, err

	case mode.IsRegular():
		fileObj := &FileObject{
			StorageType: LCL,
			Path:        path,
			Ext:         filepath.Ext(path),
			Filename:    filepath.Base(path),
		}
		files = append(files, fileObj)
	}

	return files, nil
}
