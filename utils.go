package golog

import (
	"os"
	"log"
)

var envPropName string = "GOLOG_LOGGING_LEVEL"

// SetLogLevel - set global logging level
func SetLogLevel(logLevel string) {
	if !isDebugLog(logLevel) && !isInfoLog(logLevel) && !isWarnLog(logLevel) && !isErrorLog(logLevel) {
		Error("Log level not recognised.")
	}
	_ = os.Setenv(envPropName, logLevel)
}

// GetLogLevel - get globnal logging level
func GetLogLevel() string {
	return os.Getenv(envPropName)
}

// LogToFile - print all logs in specified file
func LogToFile(filepath string, append bool) {
	logfile, err := os.OpenFile(filepath, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		Error(err)
	}

	log.SetOutput(logfile)

	defer logfile.Close()
}

func isDebugLog(logLevel string) bool {
	if logLevel == "DEBUG" || logLevel == "10" {
		return true
	}
	return false
}

// default logging level
func isInfoLog(logLevel string) bool {
	if isDebugLog(logLevel) || logLevel == "INFO" || logLevel == "20" || logLevel == "" {
		return true
	}
	return false
}

func isWarnLog(logLevel string) bool {
	if isInfoLog(logLevel) || logLevel == "WARN" || logLevel == "30" {
		return true
	}
	return false
}

func isSuccessLog(logLevel string) bool {
	if isWarnLog(logLevel) || logLevel == "SUCCESS" || logLevel == "30" {
		return true
	}
	return false
}

func isErrorLog(logLevel string) bool {
	if isWarnLog(logLevel) || logLevel == "ERROR" || logLevel == "40" {
		return true
	}
	return false
}

// SetLogFormat - set the format to [ DATE TIME LOG ]
func SetLogFormat() {
	log.SetFlags(log.Ldate | log.Ltime)
}

// SetEmptyLogFormat - unset the log format
func UnsetLogFormat() {
	log.SetFlags(0)
}

/*
To set logging levels, either use the above func SetLogLevel, or
$ export GOLOG_LOGGING_LEVEL="<log level>"
*/
