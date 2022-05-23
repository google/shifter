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

package api

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	generator "shifter/generators"
	lib "shifter/lib"
	os "shifter/openshift"
	ops "shifter/ops"
	"shifter/processor"

	"github.com/gin-gonic/gin"
)

func (server *Server) Convert(ctx *gin.Context) {
	var openshift os.Openshift

	// Create API Unique RUN ID
	//uuid := uuid.New().String()
	suid := ops.CreateSUID("")
	convert := Convert{}
	// using BindJson method to serialize body with struct
	if err := ctx.BindJSON(&convert); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var openshift os.Openshift
	openshift.Endpoint = convert.Shifter.ClusterConfig.BaseUrl
	openshift.AuthToken = convert.Shifter.ClusterConfig.BearerToken

	// Process Each Item
	for _, item := range convert.Items {
		// Confirm Project/Namespace Exists
		deploymentConfig := openshift.GetDeploymentConfig(item.Namespace.ObjectMeta.Name, item.DeploymentConfig.ObjectMeta.Name)

		u, err := json.Marshal(deploymentConfig)
		if err != nil {
			panic(err)
		}

		// Handle the Conversion of the Manifests and File Writing
		var generator generator.Generator
		var objs []lib.K8sobject
		obj := processor.Processor(u, "DeploymentConfig", nil)
		objs = append(objs, obj)
		convertedObjects := generator.Yaml(item.DeploymentConfig.ObjectMeta.Name, objs)
		for _, conObj := range convertedObjects {
			fileObj := &ops.FileObject{
				//StorageType: "GCS",
				//SourcePath:  ("gs://shifter-lz-002-sample-files/" + uuid + "/" + item.Namespace.ObjectMeta.Name + "/" + item.DeploymentConfig.ObjectMeta.Name),
				StorageType:   server.config.serverStorage.storageType,
				SourcePath:    (server.config.serverStorage.sourcePath + "/" + uuid + "/" + item.Namespace.ObjectMeta.Name + "/" + item.DeploymentConfig.ObjectMeta.Name),
				Ext:           "yaml",
				Content:       conObj.Payload,
				ContentLength: conObj.Payload.Len(),
			}
			fileObj.WriteFile()
		}
	}

	// Zip / Package Converted Objects
	err := server.PackageConversionObjects(suid)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	// Construct API Endpoint Response
	r := ResponseConvert{
		SUID:    suid,
		Message: "Converted..." + string(len(convert.Items)) + " Objects",
	}
	ctx.JSON(http.StatusOK, r)
}

func (server *Server) PackageConversionObjects(suid ops.SUID) error {

	file, err := os.Create(server.config.serverStorage.outputPath + "/" + suid.DownloadId + ".zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := zip.NewWriter(file)
	defer w.Close()

	walker := func(path string, info os.FileInfo, err error) error {
		//fmt.Printf("Crawling: %#v\n", path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Ensure that `path` is not absolute; it should not start with "/".
		// This snippet happens to work because I don't use
		// absolute paths, but ensure your real-world code
		// transforms path into a zip-root relative path.
		f, err := w.Create(path)
		if err != nil {
			return err
		}

		_, err = io.Copy(f, file)
		if err != nil {
			return err
		}
		return nil
	}
	err = filepath.Walk(server.config.serverStorage.sourcePath+"/"+suid.DirectoryName+"/", walker)
	if err != nil {
		return errors.New("Unable to resolve or find Download ID")
	}
	return nil
}
