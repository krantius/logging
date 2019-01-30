package logging

import (
	"fmt"
	"runtime"
	"strings"
)

func Trace(msg string) {
	output(TRACE, msg)
}

func Tracef(msg string, args ...interface{}) {
	output(TRACE, fmt.Sprintf(msg, args...))
}

func Debug(msg string) {
	output(DEBUG, msg)
}

func Debugf(msg string, args ...interface{}) {
	output(DEBUG, fmt.Sprintf(msg, args...))
}

func Info(msg string) {
	output(INFO, msg)
}

func Infof(msg string, args ...interface{}) {
	output(INFO, fmt.Sprintf(msg, args...))
}

func Warning(msg string) {
	output(WARNING, msg)
}

func Warningf(msg string, args ...interface{}) {
	output(WARNING, fmt.Sprintf(msg, args...))
}

func Error(msg string) {
	output(ERROR, msg)
}

func Errorf(msg string, args ...interface{}) {
	output(ERROR, fmt.Sprintf(msg, args...))
}

func output(l Level, msg string) {
	if !shouldLog(l) {
		return
	}

	_, f, line, _ := runtime.Caller(2)

	f = f[strings.LastIndex(f, "/")+1:]

	switch l {
	case TRACE:
		logger.Printf("%s:%d [TRACE] %s\n", f, line, msg)
	case DEBUG:
		logger.Printf("%s:%d [DEBUG] %s\n", f, line, msg)
	case INFO:
		logger.Printf("%s:%d [INFO] %s\n", f, line, msg)
	case WARNING:
		logger.Printf("%s:%d [WARN] %s\n", f, line, msg)
	case ERROR:
		logger.Printf("%s:%d [ERROR] %s\n", f, line, msg)
	}
}

func shouldLog(l Level) bool {
	return l >= level
}
