/*
Copyright 2019 Google LLC
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ops

import "bytes"

//type File struct {
//	Filename string `json:"filename"`
//	Link     string `json:"link"`
//}

type Converter struct {
	UUID       string // Unique ID of the Run
	InputType  string
	SourcePath string
	Generator  string
	OutputPath string
	Flags      map[string]string

	SourceFiles []*FileObject
	OutputFiles []*FileObject

	Logs []*LogObject
}

type FileObject struct {
	UUID        string // Unique ID of the Run
	StorageType string // GCS or LCL Storage
	SourcePath  string // Bucket or Local Path
	Ext         string // File Extention

	//Root     string // Root Directory "/data" LCL Storage
	//Bucket   string // Root Bucket Nmae GCS Storage
	//Path     string // Sub Path ["raw", "output"]
	Filename string // Filename

	Content       bytes.Buffer // Content as Bytes Buffer
	ContentLength int          // Content Length (len(bytes.buffer))
}

type DownloadFile struct {
	Link     string `json:"link"`
	Filename string `json:"filename"`
}

type LogObject struct {
}
