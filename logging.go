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

var traceLogger = log.New(os.Stdout, "T ", log.Ltime)
var debugLogger = log.New(os.Stdout, "D ", log.Ltime)
var infoLogger = log.New(os.Stdout, "I ", log.Ltime)
var warnLogger = log.New(os.Stdout, "W ", log.Ltime)
var errorLogger = log.New(os.Stdout, "E ", log.Ltime)

var level = INFO

// SetLevel sets the logging level to the specified value
func SetLevel(l Level) {
	if l < TRACE || l > ERROR {
		return
	}

	level = l
}
