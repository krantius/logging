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
		traceLogger.Printf("%s %s:%d\n", msg, f, line)
	case DEBUG:
		debugLogger.Printf("%s %s:%d\n", msg, f, line)
	case INFO:
		infoLogger.Printf("%s %s:%d\n", msg, f, line)
	case WARNING:
		warnLogger.Printf("%s %s:%d\n", msg, f, line)
	case ERROR:
		errorLogger.Printf("%s %s:%d\n", msg, f, line)
	}
}

func shouldLog(l Level) bool {
	return l >= level
}
