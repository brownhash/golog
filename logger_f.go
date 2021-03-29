package golog

import (
	"fmt"
	"log"
	"runtime"
)

// Debugf - log debugging messages in same line
func Debugf(message interface{}) {
	var fileInfo string = ""

	_, file, lineNum, ok := runtime.Caller(1)

	if ok {
		fileInfo = fmt.Sprintf(NoticeColor, fmt.Sprintf("%s::%d", file, lineNum))
	}

	if isDebugLog(GetLogLevel()) {
		formatter := fmt.Sprintf("%s %s", fileInfo, fmt.Sprintf(DebugColor, message))
		log.Printf(formatter)
	}
}

// Infof - log informative messages in same line
func Infof(message interface{}) {
	if isInfoLog(GetLogLevel()) {
		formatter := fmt.Sprintf(InfoColor, message)
		log.Printf(formatter)
	}
}

// Warnf - log warning messages in same line
func Warnf(message interface{}) {
	if isWarnLog(GetLogLevel()) {
		formatter := fmt.Sprintf(WarningColor, message)
		log.Printf(formatter)
	}
}

// Successf - log success messages in same line
func Successf(message interface{}) {
	if isSuccessLog(GetLogLevel()) {
		formatter := fmt.Sprintf(SuccessColor, message)
		log.Printf(formatter)
	}
}

// Errorf - log error messages in same line
func Errorf(message interface{}) {
	if isErrorLog(GetLogLevel()) {
		formatter := fmt.Sprintf(ErrorColor, message)
		log.Fatalf(formatter)
	}
}
