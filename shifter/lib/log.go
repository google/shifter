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
		log.Println(colourCyan + "🗨️ [" + level + "] " + colourReset + message)
	case "info":
		log.Println(colourGreen + "💡[" + level + "] " + colourReset + message)
	case "warn":
		log.Println(colourYellow + "⚠️ [" + level + "] " + colourReset + message)
	case "error":
		log.Println(colourRed + "❌[" + level + "] " + colourReset + message + " - " + err.Error())
	case "fatal":
		log.Println(colourRed + "☠️ [" + level + "] " + colourReset + message + " - " + err.Error())
	case "panic":
		log.Printf("☠️ [%s] %s", level, message)
		panic(err)
	}
}
