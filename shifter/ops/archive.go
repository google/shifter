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
	"archive/zip"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"shifter/lib"
)

func Archive(sourcePath string, outputPath string, suuid SUUID) error {

	lib.CLog("info", "Creating archive file of the results")
	if _, err := os.Stat(outputPath + "/"); os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(outputPath+"/"), 0700)
		if err != nil {
			lib.CLog("error", "Creating archive directories", err)
			return (err)
		}
	}

	file, err := os.Create(outputPath + "/" + suuid.DownloadId + ".zip")
	if err != nil {
		lib.CLog("error", "Creating archive", err)
		return (err)
	}
	defer file.Close() // TODO - Fix gosec error

	w := zip.NewWriter(file)
	defer w.Close() // TODO - Fix gosec error

	walker := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Error: Traversing Directory
			log.Printf("🧰 ❌ ERROR: Creating Archive File.")
			return err
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			// Error: Opening File
			log.Printf("🧰 ❌ ERROR: Opening Archive File.")
			return err
		}
		defer file.Close() // TODO - Fix gosec error

		// Ensure that `path` is not absolute; it should not start with "/".
		// This snippet happens to work because I don't use
		// absolute paths, but ensure your real-world code
		// transforms path into a zip-root relative path.
		f, err := w.Create(path)
		if err != nil {
			// Error: Creating File within Archive
			log.Printf("🧰 ❌ ERROR: Creating File within Archive.")
			return err
		}

		_, err = io.Copy(f, file)
		if err != nil {
			// Error: Creating Coping Buffer content to Archive File
			log.Printf("🧰 ❌ ERROR: Writing Archive File.")
			return err
		}
		return nil
	}
	err = filepath.Walk(sourcePath+"/"+suuid.DirectoryName+"/", walker)
	if err != nil {
		// Error: Unable to resolve or find Download ID
		log.Printf("🧰 ❌ ERROR: Unable to resolve or find Download ID.")
		return errors.New("Unable to resolve or find Download ID")
	}
	return nil
}
