package logging

import "fmt"

// Trace
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
		logger.Printf("TRACE %s\n", msg)
	case DEBUG:
		logger.Printf("DEBUG %s\n", msg)
	case INFO:
		logger.Printf("INFO %s\n", msg)
	case WARNING:
		logger.Printf("WARN %s\n", msg)
	case ERROR:
		logger.Printf("ERROR %s\n", msg)
	}
}

func shouldLog(l Level) bool {
	return l >= level
}
