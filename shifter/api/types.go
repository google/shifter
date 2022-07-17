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

package api

import ops "shifter/ops"

type Shifter struct {
	ClusterConfig *ClusterConfig `json:"clusterConfig"`
}

type ClusterConfig struct {
	ConnectionName string `json:"connectionName"`
	BaseUrl        string `json:"baseUrl"`
	BearerToken    string `json:"bearerToken"`
	Username       string `json:"username"`
	Password       string `json:"password"`
}

type Downloads struct {
	Items *[]Download `json:"items"`
}

type Download struct {
	Link        string `json:"link"`
	Name        string `json:"name"`
	Uuid        string `json:"uuid"`
	DisplayName string `json:"displayName"`
}

type ResponseDownload struct {
	SUID    ops.SUID `json:"suid"`
	Message string   `json:"message"`
}

type ResponseDownloads struct {
	Items   []*ops.SUID `json:"items"`
	Message string      `json:"message"`
}

type ServerConfig struct {
	serverAddress   string
	serverPort      string
	storagePlatform string
	//gcsBucket       string
	serverStorage ServerStorage
}

type ServerStorage struct {
	description string
	storageType string
	sourcePath  string
	outputPath  string
}
