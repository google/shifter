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
	"context"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

// TODO - Error Handling
// Takes GCS Bucket Path and Returns the Bucket & File Path
func GCSPathDeconstruction(gscPath string) (string, string) {
	route := strings.SplitN(strings.ReplaceAll(gscPath, "gs://", ""), "/", 2)
	bucket := route[0]
	path := ""
	// if Path contains
	if len(route) >= 2 {
		path = route[1]
	}
	return bucket, path
}

// Write bytes.Buffer to GSC Bucket as File
func (fileObj *FileObject) WriteGCSFile() error {
	log.Printf("🧰 💡 INFO: Writing Shifter File Object to GCS")

	bucket, prefix := GCSPathDeconstruction(fileObj.Path)

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		// Error: Connecting to Provided GCS Bucket
		log.Printf("🧰 ❌ ERROR: Connecting to Provided GCS Bucket: '%s'.", fileObj.Path)
		return err
	} else {
		// Success: Connecting to Provided GCS Bucket
		log.Printf("🧰 ✅ SUCCESS: Connecting to Provided GCS Bucket: '%s'.", fileObj.Path)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*60)
	defer cancel()

	wc := client.Bucket(bucket).Object(prefix).NewWriter(ctx)
	wc.ChunkSize = 0

	if _, err = io.Copy(wc, &fileObj.Content); err != nil {
		// Error: Writing to GCS Bucket
		log.Printf("🧰 ❌ ERROR: Writing Content to GCS Bucket: '%s'.", fileObj.Path)
		return err
	} else {
		// Success: Writing to GCS Bucket
		log.Printf("🧰 ✅ SUCCESS: Writing Content to GCS Bucket: '%s'.", fileObj.Path)
	}

	if err := wc.Close(); err != nil {
		// Error: Closing GCS Bucket Writer
		log.Printf("🧰 ❌ ERROR: Unable to close GCS Bucket Writer.")
		return err
	} else {
		// Success: GCS Writer Closed
		log.Printf("🧰 ✅ SUCCESS: Closing GCS Bucket Writer: '%s'.", fileObj.Path)
	}

	// Successfull Writen file to GCS Bucket
	log.Printf("🧰 ✅ SUCCESS: File written to GCS Bucket: '%s'.", fileObj.Path)
	return nil
}

// Load GSC File from Bucket to bytes.Buffer
func (fileObj *FileObject) LoadGCSFile() error {
	log.Printf("🧰 💡 INFO: Loading Shifter File Object from GCS")

	bucket, prefix := GCSPathDeconstruction(fileObj.Path)

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		// Error: Connecting to Provided GCS Bucket
		log.Printf("🧰 ❌ ERROR: Connecting to Provided GCS Bucket File: '%s'.", fileObj.Path)
		return err
	} else {
		// Success: Connecting to Provided GCS Bucket
		log.Printf("🧰 ✅ SUCCESS: Connecting to Provided GCS Bucket: '%s'.", fileObj.Path)
	}

	// Configure Client to Bucket Path
	bkt := client.Bucket(bucket)

	rc, err := bkt.Object(prefix).NewReader(ctx)
	if err != nil {
		// Error: Unable to open file in GCS Bucket
		log.Printf("🧰 ❌ ERROR: Unable to open file in GCS Bucket:'%s'", bucket)
		return err
	} else {
		// Success: Opening file in GCS Bucket
		log.Printf("🧰 ✅ SUCCESS: Opening file in GCS Bucket: '%s'.", fileObj.Path)
	}
	defer rc.Close()

	// Attempt to Read File Contents to buffer.
	r, err := bkt.Object(prefix).NewReader(ctx)
	if err != nil {
		// ERROR: Reading GCS file contents into Bytes Buffer
		log.Printf("🧰 ❌ ERROR: Reading GCS file contents into Buffer.")
		return err
	} else {
		// SCUCCESS: Reading Local File System File
		log.Printf("🧰 ✅ SUCCESS: Reading Shifter File Object from Local Filesystem.")
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		// ERROR: Reading file contents into Buffer
		log.Printf("🧰 ❌ ERROR: Reading file contents into Buffer.")
		return err
	} else {
		// Success: Reading file contents into Buffer
		log.Printf("🧰 ✅ SUCCESS: Reading file contents into Buffer.")
	}

	fileObj.Content = *bytes.NewBuffer(data)
	fileObj.ContentLength = len(data)

	// SCUCCESS: Reading Local File System File
	log.Printf("🧰 ✅ SUCCESS: Reading Shifter File Object from GCS Bucket.")
	return nil
}

// Walk a GCS Path and Retrieve a list og Files
func ProcessGCSPath(path string) ([]*FileObject, error) {

	bucket, prefix := GCSPathDeconstruction(path)

	var files []*FileObject

	// Conect to cloud Storage
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		// Error: Connecting to Provided GCS Bucket
		log.Printf("🧰 ❌ ERROR: Connecting to Provided GCS Bucket: '%s'.", bucket)
		return files, err
	} else {
		// Success: Connecting to Provided GCS Bucket
		log.Printf("🧰 ✅ SUCCESS: Connecting to Provided GCS Bucket: '%s'.", bucket)
	}

	// Configure Client to Bucket Path
	bkt := client.Bucket(bucket)
	// Filter the Bucket Objects by prefix
	query := &storage.Query{Prefix: prefix}
	it := bkt.Objects(ctx, query)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// Error: Filtering GCS Bucket Contents
			log.Printf("🧰 ❌ ERROR: Filtering GCS Bucket Contents.")
			return files, err
		} else {
			// Success: Filtering GCS Bucket Contents
			log.Printf("🧰 ✅ SUCCESS: Filtering GCS Bucket Contents")
		}

		// Create File Object for every file in Bucket
		fileObj := &FileObject{
			StorageType: GCS,
			Path:        ("gs://" + bucket + "/" + attrs.Name),
			Ext:         filepath.Ext(attrs.Name),
		}
		// Add File Object to Array of Files
		files = append(files, fileObj)
	}

	// Success Processing GCS File Path
	return files, nil
}
