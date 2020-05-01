package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	TRACE   = "TRACE"
	DEBUG   = "DEBUG"
	INFO    = "INFO"
	WARNING = "WARNING"
	ERROR   = "ERROR"
)

var LogLevelsMap = map[string]int{
	"TRACE":   0,
	"DEBUG":   1,
	"INFO":    2,
	"WARNING": 3,
	"ERROR":   4,
}

var logLevel = 2

var (
	traceLogger   = log.New(os.Stdout, "[TRACE]: ", log.Ldate|log.Ltime|log.Llongfile)
	debugLogger   = log.New(os.Stdout, "[DEBUG]: ", log.Ldate|log.Ltime|log.Llongfile)
	infoLogger    = log.New(os.Stdout, "[INFO]: ", log.Ldate|log.Ltime|log.Llongfile)
	warningLogger = log.New(os.Stdout, "[WARNING]: ", log.Ldate|log.Ltime|log.Llongfile)
	errorLogger   = log.New(os.Stderr, "[ERROR]: ", log.Ldate|log.Ltime|log.Llongfile)
)

func Trace(v ...interface{}) {
	if logLevel == LogLevelsMap[TRACE] {
		traceLogger.Output(2, fmt.Sprintln(v...))
	}
}

func Tracef(format string, v ...interface{}) {
	if logLevel == LogLevelsMap[TRACE] {
		traceLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

func Debug(v ...interface{}) {
	if logLevel >= LogLevelsMap[TRACE] && logLevel <= LogLevelsMap[DEBUG] {
		debugLogger.Output(2, fmt.Sprintln(v...))
	}
}

func Debugf(format string, v ...interface{}) {
	if logLevel >= LogLevelsMap[TRACE] && logLevel <= LogLevelsMap[DEBUG] {
		debugLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

func Info(v ...interface{}) {
	if logLevel >= LogLevelsMap[TRACE] && logLevel <= LogLevelsMap[INFO] {
		debugLogger.Output(2, fmt.Sprintln(v...))
	}
}

func Infof(format string, v ...interface{}) {
	if logLevel >= LogLevelsMap[TRACE] && logLevel <= LogLevelsMap[INFO] {
		debugLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

func Warn(v ...interface{}) {
	if logLevel >= LogLevelsMap[TRACE] && logLevel <= LogLevelsMap[WARNING] {
		debugLogger.Output(2, fmt.Sprintln(v...))
	}
}

func Warnf(format string, v ...interface{}) {
	if logLevel >= LogLevelsMap[TRACE] && logLevel <= LogLevelsMap[WARNING] {
		debugLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

func Error(v ...interface{}) {
	if logLevel >= LogLevelsMap[TRACE] && logLevel <= LogLevelsMap[ERROR] {
		debugLogger.Output(2, fmt.Sprintln(v...))
	}
}

func Errorf(format string, v ...interface{}) {
	if logLevel >= LogLevelsMap[TRACE] && logLevel <= LogLevelsMap[ERROR] {
		debugLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

func SetLogLevel(loggingLevel string) {
	logLevel = LogLevelsMap[loggingLevel]
}
