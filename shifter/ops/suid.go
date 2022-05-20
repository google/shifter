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
	"strings"
	"time"

	"github.com/google/uuid"
)

//Used for String Splitting Pre and Post Hash
const seperator string = "#!#"

type SUID struct {
<<<<<<< HEAD
	UUID          string    `json:"link"`
	Name          string    `json:"name"`
	TimeStamp     time.Time `json:"timestamp"`
	DisplayName   string    `json:"displayName"`
	DirectoryName string    `json:"directoryName"`
	DownloadId    string    `json:"downloadId"`
	// Private
	longname string
	nameHash string
=======
	UUID        string    `json:"link"`
	Name        string    `json:"name"`
	TimeStamp   time.Time `json:"timestamp"`
	DisplayName string    `json:"displayName"`
	// Private
	longname    string
	nameHash    string
	OutLongname string
>>>>>>> 85e899b (Landed the encoding of conversion objects metadata into foldername)
}

func (s *SUID) name() string {
	return s.Name
}

func (s *SUID) hash() {
	s.nameHash = base64.StdEncoding.EncodeToString([]byte(s.longname))
}

func (s *SUID) Meta() {
	fmt.Println("_+_+_+_+_+_+_+_+_+_+_+_+_+_")
	fmt.Println("_+_+_+_+_+_+_+_+_+_+_+_+_+_")
	fmt.Println(s.UUID)
	fmt.Println(s.Name)
	fmt.Println(s.TimeStamp)
	fmt.Println(s.DisplayName)
	fmt.Println(s.longname)
	fmt.Println(s.nameHash)
	fmt.Println("_+_+_+_+_+_+_+_+_+_+_+_+_+_")
	fmt.Println("_+_+_+_+_+_+_+_+_+_+_+_+_+_")
	fmt.Println("")
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
<<<<<<< HEAD
	suid.DirectoryName = suid.nameHash
	suid.DownloadId = suid.nameHash
	suid.DisplayName = fmt.Sprintf("%s - %s", suid.TimeStamp.Format(time.RFC1123), suid.Name)
	return suid
}

func ResolveSUID(downloadId string) (SUID, error) {
	// Create New SUID Object
	suid := SUID{}
	if downloadId == "" {
		return suid, errors.New("Download ID or Filename Hash must be provided when Resolving SUID")
	}
	suid.nameHash = downloadId
	decoded, err := base64.StdEncoding.DecodeString(suid.nameHash)
=======

	suid.OutLongname = suid.nameHash
	return suid
}

func ResolveSUID(hash string) (SUID, error) {
	// Create New SUID Object
	suid := SUID{}
	if hash == "" {
		return suid, errors.New("Filename Hash must be provided when Resolving SUID")
	}
	suid.nameHash = hash
	decoded, err := base64.StdEncoding.DecodeString(hash)
>>>>>>> 85e899b (Landed the encoding of conversion objects metadata into foldername)
	if err != nil {
		return suid, err
	}
	suid.longname = string(decoded)
	items := strings.Split(suid.longname, seperator)
	t, err := time.Parse(time.RFC1123, items[0])
	if err != nil {
		return suid, err
	}
	suid.TimeStamp = t
	suid.UUID = items[1]
	suid.Name = items[2]
<<<<<<< HEAD
	suid.DisplayName = fmt.Sprintf("%s - %s", suid.TimeStamp.Format(time.RFC1123), suid.Name)
	suid.DirectoryName = suid.nameHash
	suid.DownloadId = suid.DirectoryName
=======

>>>>>>> 85e899b (Landed the encoding of conversion objects metadata into foldername)
	return suid, nil

}
