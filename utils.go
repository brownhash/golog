package golog

import (
	"os"
)

// SetLogLevel - set global logging level
func SetLogLevel(logLevel string) {
	if (logLevel != "DEBUG" && logLevel != "10") && (logLevel != "INFO" && logLevel != "20") && (logLevel != "WARN" && logLevel != "30") && (logLevel != "ERROR" && logLevel != "40") {
		Error("Log level not recognised.")
	}
	_ = os.Setenv("GOLOG_LOGGING_LEVEL", logLevel)
}

func GetLogLevel() string {
	return os.Getenv("GOLOG_LOGGING_LEVEL")
}

func isDebugLog(logLevel string) bool {
	if logLevel == "DEBUG" || logLevel == "10" || logLevel == "" {
		return true
	}
	return false
}

func isInfoLog(logLevel string) bool {
	if isDebugLog(logLevel) || logLevel == "INFO" || logLevel == "20" {
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
