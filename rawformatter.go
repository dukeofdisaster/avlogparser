package avlogparser

import (
	"log"
	"strings"
)

// Names : array of fixed log names
var Names = []string{
	"timestamp",
	"dest_ip",
	"dest_port",
	"src_ip",
	"src_port",
	"http_user_agent",
	"url",
	"payload",
}

// ValueExtract : takes given string and a given key then extracts the value of the key.
// Functionality is dependent on the format of the log.
// DOD NOTE: Name convention and comment style required.
func ValueExtract(myString string, key string) string {
	keyLength := len(key)
	keyStart := strings.Index(myString, key)
	valStart := keyLength + keyStart + 4
	if keyStart == -1 {
		return "none"
	}
	if key == "dest_port" {
		dstart := keyStart + 12
		// NOTE: slicing based on max port number could/should potentially
		// improve performance on larger data sets since we don't have to perform the op on
		// the entire rest of the string, but whatever.
		return myString[(dstart) : dstart+strings.Index(myString[dstart:], ",")]
	}
	if key == "src_port" {
		sstart := keyStart + 11
		return myString[(sstart) : sstart+strings.Index(myString[sstart:], ",")]
	}
	valEnd := valStart + strings.Index(myString[valStart:], "\"")
	return myString[valStart:valEnd]
}

// ExtractAll : loops through the local array of prefixed key names and extracts
// all the data from a given string (i.e from a log line)
// Could also overload to take myString and myAarray for mroe flexibility
func ExtractAll(myString string) []string {
	var goods []string
	for i := range Names {
		goods = append(goods, ValueExtract(myString, Names[i]))
	}
	return goods
}

// CheckError : determines if the given error is nil and emits a given message
func CheckError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
