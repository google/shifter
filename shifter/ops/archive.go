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
	"archive/zip"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
)

// TODO - Remove Function?
/*
func Archive(srcPath string, fileName string) error {
	// 1. Create a ZIP file and zip.Writer
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()

	// 2. Go through all the files of the source
	return filepath.Walk(srcPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 3. Create a local file header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// set compression
		header.Method = zip.Deflate

		// 4. Set relative path of a file as the header name
		header.Name, err = filepath.Rel(filepath.Dir(srcPath), path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += "/"
		}

		// 5. Create writer for the file header and save content of the file
		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(headerWriter, f)
		return err
	})
}*/

func Archive(sourcePath string, outputPath string, suid SUID) error {

	if _, err := os.Stat(outputPath + "/"); os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(outputPath+"/"), 0600) // Create output directory
		if err != nil {
			// Error: Creating Directories in Archive
			log.Printf("🧰 ❌ ERROR: Creating Directories in Archive.")
			return (err)
		}
	}
	file, err := os.Create(outputPath + "/" + suid.DownloadId + ".zip")
	if err != nil {
		// Error: Creating Archive File
		log.Printf("🧰 ❌ ERROR: Creating Archive File.")
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
	err = filepath.Walk(sourcePath+"/"+suid.DirectoryName+"/", walker)
	if err != nil {
		// Error: Unable to resolve or find Download ID
		log.Printf("🧰 ❌ ERROR: Unable to resolve or find Download ID.")
		return errors.New("Unable to resolve or find Download ID")
	}
	return nil
}
