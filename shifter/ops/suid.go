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
	UUID        string    `json:"link"`
	Name        string    `json:"name"`
	TimeStamp   time.Time `json:"timestamp"`
	DisplayName string    `json:"displayName"`
	// Private
	longname    string
	nameHash    string
	OutLongname string
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

	return suid, nil

}
