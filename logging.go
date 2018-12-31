package logging

import (
	"log"
	"os"
)

// Level defines the logging level
type Level int

// Different log levels starting at most verbose going to least verbose
const (
	TRACE Level = iota
	DEBUG
	INFO
	WARNING
	ERROR
)

var logger = log.New(os.Stdout, "", 19)
var level = INFO

// SetLevel sets the logging level to the specified value
func SetLevel(l Level) {
	if l < TRACE || l > ERROR {
		return
	}

	level = l
}
