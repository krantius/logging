package logging

import "fmt"

func Trace(msg string) {
	output(TRACE, msg)
}

func Tracef(msg string, args ...interface{}) {
	tmp := make([]interface{}, len(args))

	for i := range args {
		tmp[i] = args[i]
	}

	Trace(fmt.Sprintf(msg, tmp...))
}

func Debug(msg string) {
	output(DEBUG, msg)
}

func Debugf(msg string, args ...interface{}) {
	tmp := make([]interface{}, len(args))

	for i := range args {
		tmp[i] = args[i]
	}

	Debug(fmt.Sprintf(msg, tmp...))
}

func Info(msg string) {
	output(INFO, msg)
}

func Infof(msg string, args ...interface{}) {
	tmp := make([]interface{}, len(args))

	for i := range args {
		tmp[i] = args[i]
	}

	Info(fmt.Sprintf(msg, tmp...))
}

func Warning(msg string) {
	output(WARNING, msg)
}

func Warningf(msg string, args ...interface{}) {
	tmp := make([]interface{}, len(args))

	for i := range args {
		tmp[i] = args[i]
	}

	Warning(fmt.Sprintf(msg, tmp...))
}

func Error(msg string) {
	output(ERROR, msg)
}

func Errorf(msg string, args ...interface{}) {
	tmp := make([]interface{}, len(args))

	for i := range args {
		tmp[i] = args[i]
	}

	Error(fmt.Sprintf(msg, tmp...))
}

func output(l Level, msg string) {
	if !shouldLog(l) {
		return
	}

	switch l {
	case TRACE:
		traceLogger.Println(msg)
	case DEBUG:
		debugLogger.Println(msg)
	case INFO:
		infoLogger.Println(msg)
	case WARNING:
		warningLogger.Println(msg)
	case ERROR:
		errorLogger.Println(msg)
	}
}

func shouldLog(l Level) bool {
	return l >= level
}
