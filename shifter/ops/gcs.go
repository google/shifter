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
func (fileObj *FileObject) WriteGCSFile() {

	bucket, prefix := GCSPathDeconstruction(fileObj.Path)

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Println("Error connecting to GCS")
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*60)
	defer cancel()

	wc := client.Bucket(bucket).Object(prefix).NewWriter(ctx)
	wc.ChunkSize = 0

	if _, err = io.Copy(wc, &fileObj.Content); err != nil {
		log.Println("io.Copy: ", err)
	}

	if err := wc.Close(); err != nil {
		log.Println("Writer.Close: ", err)
	}
}

// Load GSC File from Bucket to bytes.Buffer
func (fileObj *FileObject) LoadGCSFile() {

	bucket, prefix := GCSPathDeconstruction(fileObj.Path)

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Configure Client to Bucket Path
	bkt := client.Bucket(bucket)

	rc, err := bkt.Object(prefix).NewReader(ctx)
	if err != nil {
		log.Println("readFile: unable to open file from bucket")
		return
	}
	defer rc.Close()

	r, err := bkt.Object(prefix).NewReader(ctx)
	if err != nil {
		log.Println(err)
		return
		// TODO: Handle error.
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		log.Println(err)
		return
	}
	fileObj.Content = *bytes.NewBuffer(data)
	fileObj.ContentLength = len(data)
}

// Walk a GCS Path and Retrieve a list og Files
func ProcessGCSPath(path string) ([]*FileObject, error) {

	bucket, prefix := GCSPathDeconstruction(path)

	var files []*FileObject

	// Conect to cloud Storage
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
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
			log.Fatal(err)
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

	return files, nil
}
