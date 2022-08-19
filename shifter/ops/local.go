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
	"os"
	"path/filepath"
	"shifter/lib"
)

// Write bytes.Buffer to Local File System
func (fileObj *FileObject) WriteLCLFile() error {
	lib.CLog("debug", "Writing file object to local file system")

	// Check if Output Directory Exits, IF Not, Make output Directory
	lib.CLog("debug", "Path: "+filepath.Dir(fileObj.Path))
	if _, err := os.Stat(filepath.Dir(fileObj.Path)); os.IsNotExist(err) {
		lib.CLog("info", "Output directory does not exist... creating")
		err := os.MkdirAll(filepath.Dir(fileObj.Path), 0700)
		if err != nil {
			lib.CLog("error", "Unable to create output directory", err)
			return err
		}
	}

	fileName := fileObj.Path + "." + fileObj.Ext
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
