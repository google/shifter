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
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"shifter/lib"
	"strings"
	"time"
)

// Used for String Splitting Pre and Post Hash
const seperator string = "#!#"

type SUUID struct {
	UUID          string    `json:"link"`
	Name          string    `json:"name"`
	TimeStamp     time.Time `json:"timestamp"`
	DisplayName   string    `json:"displayName"`
	DirectoryName string    `json:"directoryName"`
	DownloadId    string    `json:"downloadId"`
	// Private
	longname string
	nameHash string
}

func (s *SUUID) name() string {
	return s.Name
}

func (s *SUUID) hash() {
	s.nameHash = base64.StdEncoding.EncodeToString([]byte(s.longname))
}

func (suuid *SUUID) Meta() {
	log.Printf("UUID: %s, \nName: %s, \nTimeStamp: %s, \nDisplayName: %s, \nLongName: %s, Hash: %s",
		suuid.UUID,
		suuid.Name,
		suuid.TimeStamp,
		suuid.DisplayName,
		suuid.nameHash,
	)
}

func CreateUUID(customName string) SUUID {
	suuid := SUUID{}
	suuid.UUID = uuid.New().String()
	suuid.TimeStamp = time.Now()

	if customName == "" {
		suuid.Name = "Shifter Conversion"
	} else {
		// TODO Clean Name Here
		suuid.Name = customName
	}

	// String Format - (Timestamp + UUID + Custom Name)
	suuid.longname = fmt.Sprintf("%s%s%s%s%s", suuid.TimeStamp.Format(time.RFC1123), seperator, suuid.UUID, seperator, suuid.Name)
	suuid.hash()
	suuid.DirectoryName = suuid.nameHash
	suuid.DownloadId = suuid.nameHash
	suuid.DisplayName = fmt.Sprintf("%s - %s", suuid.TimeStamp.Format(time.RFC1123), suuid.Name)

	return suuid
}

func ResolveUUID(downloadId string) (SUUID, error) {
	suuid := SUUID{}

	if downloadId == "" {
		return suuid, errors.New("Download ID or Filename Hash must be provided when Resolving UUID")
	}
	suuid.nameHash = downloadId

	decoded, err := base64.StdEncoding.DecodeString(suuid.nameHash)
	if err != nil {
		lib.CLog("error", "Decoding UUID Hash", err)
		return suuid, err
	}
	suuid.longname = string(decoded)

	items := strings.Split(suuid.longname, seperator)
	t, err := time.Parse(time.RFC1123, items[0])
	if err != nil {
		lib.CLog("error", "Parsing UUID Timestamp", err)
		return suuid, err
	}

	suuid.TimeStamp = t
	suuid.UUID = items[1]
	suuid.Name = items[2]
	suuid.DisplayName = fmt.Sprintf("%s - %s", suuid.TimeStamp.Format(time.RFC1123), suuid.Name)
	suuid.DirectoryName = suuid.nameHash
	suuid.DownloadId = suuid.DirectoryName

	return suuid, nil
}
