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
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"io"
	"log"
	"time"
)

// Steam file upload to Google Cloud Storage

func StreamFileUpload(w io.Writer, bucket, object string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("Error connecting to GCS: %v", err)
	}
	defer client.Close()

	r := io.Reader(w)

	b := []byte(r)
	buf := bytes.NewBuffer(b)

	ctx, cancel := context.WithTimeout(ctx, time.Second*60)
	defer cancel()

	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	wc.ChunkSize = 0

	if _, err = io.Copy(wc, buf); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}

	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}
	log.Println(w, "%v uploaded to %v.\n", object, bucket)

	return nil
}
