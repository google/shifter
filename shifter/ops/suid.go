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
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Used for String Splitting Pre and Post Hash
const seperator string = "#!#"

type SUID struct {
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

func (s *SUID) name() string {
	return s.Name
}

func (s *SUID) hash() {
	// Constructing the SUID HASH
	s.nameHash = base64.StdEncoding.EncodeToString([]byte(s.longname))
}

// Display SUID Object Content for Debugging.
func (suid *SUID) Meta() {
	// Log SUID Object Data.
	log.Printf("🌐 📜 INFO: Shifter SUID Object \n[\nUUID: %s, \nName: %s, \nTimeStamp: %s, \nDisplay Name: %s, \nLong Name: %s, \nName HASH: %s \n]",
		suid.UUID,
		suid.Name,
		suid.TimeStamp,
		suid.DisplayName,
		suid.nameHash,
	)
}

func CreateSUID(customName string) SUID {
	// Create SUID Object
	suid := SUID{}
	// Assign Variables for Uniqueness
	suid.UUID = uuid.New().String()
	suid.TimeStamp = time.Now()

	// Ensure We always have a Conversion Name
	if customName == "" {
		// Configure Custom Name
		suid.Name = "Shifter Conversion"
	} else {
		// TODO Clean Name Here
		suid.Name = customName
	}

	// String Format - (Timestamp + UUID + Custom Name)
	suid.longname = fmt.Sprintf("%s%s%s%s%s", suid.TimeStamp.Format(time.RFC1123),
		seperator, suid.UUID, seperator, suid.Name)

	suid.hash()
	suid.DirectoryName = suid.nameHash
	suid.DownloadId = suid.nameHash
	suid.DisplayName = fmt.Sprintf("%s - %s", suid.TimeStamp.Format(time.RFC1123), suid.Name)

	// SUID Created, Return Return SUID
	return suid
}

func ResolveSUID(downloadId string) (SUID, error) {
	// Create New SUID Object
	suid := SUID{}
	if downloadId == "" {
		// Error - Download ID is NUll
		return suid, errors.New("🌐 ❌ ERROR: Download ID or Filename Hash must be provided when Resolving SUID")
	} else {
		// Log Shifter SUID Name Hash
		log.Printf("🌐 📜 INFO: Set Shifter SUID Name")
		suid.nameHash = downloadId
	}

	decoded, err := base64.StdEncoding.DecodeString(suid.nameHash)
	if err != nil {
		// Error: Unable Decode Name Hash
		log.Printf("🌐 ❌ ERROR: Unable to Decode SUID Hash")
		return suid, err
	} else {
		// Log Shifter SUID Long Name
		log.Printf("🌐 📜 INFO: Set Shifter SUID Long Name")
		suid.longname = string(decoded)
	}

	// Split the SUID Long Name by Seperator
	items := strings.Split(suid.longname, seperator)
	t, err := time.Parse(time.RFC1123, items[0])
	if err != nil {
		log.Printf("🌐 ❌ ERROR: Unable Parse Shifter SUID TimeStamp")
		return suid, err
	}
	suid.TimeStamp = t
	suid.UUID = items[1]
	suid.Name = items[2]
	suid.DisplayName = fmt.Sprintf("%s - %s", suid.TimeStamp.Format(time.RFC1123), suid.Name)
	suid.DirectoryName = suid.nameHash
	suid.DownloadId = suid.DirectoryName

	// SUID Resolved, Return
	return suid, nil
}
