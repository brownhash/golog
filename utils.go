package golog

import (
	"os"
)

var envPropName string = "GOLOG_LOGGING_LEVEL"

// SetLogLevel - set global logging level
func SetLogLevel(logLevel string) {
	if !isDebugLog(logLevel) && !isInfoLog(logLevel) && !isWarnLog(logLevel) && !isErrorLog(logLevel) {
		Error("Log level not recognised.")
	}
	_ = os.Setenv(envPropName, logLevel)
}

func GetLogLevel() string {
	return os.Getenv(envPropName)
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

/*
To set logging levels, either use the above func SetLogLevel, or
$ export GOLOG_LOGGING_LEVEL="<log level>"
*/
