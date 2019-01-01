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

var traceLogger = log.New(os.Stdout, "TRACE ", 19)
var debugLogger = log.New(os.Stdout, "DEBUG ", 19)
var infoLogger = log.New(os.Stdout, "INFO ", 19)
var warningLogger = log.New(os.Stdout, "WARNING ", 19)
var errorLogger = log.New(os.Stdout, "ERROR ", 19)

var level = INFO

// SetLevel sets the logging level to the specified value
func SetLevel(l Level) {
	if l < TRACE || l > ERROR {
		return
	}

	level = l
}
