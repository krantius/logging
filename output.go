package logging

import "fmt"

func Trace(msg string) {
	output(TRACE, msg)
}

func Tracef(msg string, args ...interface{}) {
	Trace(fmt.Sprintf(msg, args...))
}

func Debug(msg string) {
	output(DEBUG, msg)
}

func Debugf(msg string, args ...interface{}) {
	Debug(fmt.Sprintf(msg, args...))
}

func Info(msg string) {
	output(INFO, msg)
}

func Infof(msg string, args ...interface{}) {
	Info(fmt.Sprintf(msg, args...))
}

func Warning(msg string) {
	output(WARNING, msg)
}

func Warningf(msg string, args ...interface{}) {
	Warning(fmt.Sprintf(msg, args...))
}

func Error(msg string) {
	output(ERROR, msg)
}

func Errorf(msg string, args ...interface{}) {
	Error(fmt.Sprintf(msg, args...))
}

func output(l Level, msg string) {
	if !shouldLog(l) {
		return
	}

	switch l {
	case TRACE:
		logger.Printf("[TRACE] %s\n", msg)
	case DEBUG:
		logger.Printf("[DEBUG] %s\n", msg)
	case INFO:
		logger.Printf("[INFO] %s\n", msg)
	case WARNING:
		logger.Printf("[WARN] %s\n", msg)
	case ERROR:
		logger.Printf("[ERROR] %s\n", msg)
	}
}

func shouldLog(l Level) bool {
	return l >= level
}
