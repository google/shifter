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

package lib

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

// Steam file upload to Local Storage
func LocalStreamFileUpload(b bytes.Buffer, directory string, object string) error {
	
	// Create New File
	f, err := os.Create(fmt.Sprintf("%s/%s", directory, object))
	if err != nil {
		log.Println(err)
		return err
	}
	defer f.Close()

	// Write the Bytes Buffer to File
	n, err := f.Write(b.Bytes())
	if err != nil {
		log.Println(err)
		return err
	}
	
	log.Println(fmt.Sprintf("%d Bytes Written to Local Storage File %s/%s", n, directory, object))

	f.Sync()

	return nil
}
