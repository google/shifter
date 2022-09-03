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

package lib

import (
	"fmt"
	"log"
	"strings"
)

const (
	colourReset  string = "\033[0m"
	colourRed    string = "\033[31m"
	colourGreen  string = "\033[32m"
	colourYellow string = "\033[33m"
	colourBlue   string = "\033[34m"
	colourPurple string = "\033[35m"
	colourCyan   string = "\033[36m"
)

var (
	message string
	err     error
)

func CLog(level string, args ...interface{}) {
	level = strings.ToUpper(level)

	message = fmt.Sprintf("%v", args[0])

	if len(args) == 2 {
		err = args[1].(error)
	} else {
		err = nil
	}

	switch strings.ToLower(level) {
	case "debug":
		log.Println(colourCyan + "🗨️ | " + level + " | " + colourReset + message)
	case "info":
		log.Println(colourGreen + "💡| " + level + " | " + colourReset + message)
	case "warn":
		log.Println(colourYellow + "⚠️ | " + level + " | " + colourReset + message)
	case "error":
		log.Println(colourRed + "❌| " + level + " | " + colourReset + message + " - " + err.Error())
	case "fatal":
		log.Println(colourRed + "☠️ | " + level + " | " + colourReset + message + " - " + err.Error())
	case "panic":
		log.Printf("☠️ [%s] %s", level, message)
		panic(err)
	}
}
