package golog

import (
	"fmt"
	"log"
	"runtime"
)

// Debug - log debugging messages
func Debug(message interface{}) {
	SetLogFormat()
	var fileInfo string = ""

	_, file, lineNum, ok := runtime.Caller(1)

	if ok {
		fileInfo = fmt.Sprintf(NoticeColor, fmt.Sprintf("%s::%d", file, lineNum))
	}

	if isDebugLog(GetLogLevel()) {
		formatter := fmt.Sprintf("%s %s", fileInfo, fmt.Sprintf(DebugColor, message))
		log.Println(formatter)
	}
}

// Info - log informative messages
func Info(message interface{}) {
	SetLogFormat()
	if isInfoLog(GetLogLevel()) {
		formatter := fmt.Sprintf(InfoColor, message)
		log.Println(formatter)
	}
}

// Warn - log warning messages
func Warn(message interface{}) {
	SetLogFormat()
	if isWarnLog(GetLogLevel()) {
		formatter := fmt.Sprintf(WarningColor, message)
		log.Println(formatter)
	}
}

// Success - log success messages
func Success(message interface{}) {
	SetLogFormat()
	if isSuccessLog(GetLogLevel()) {
		formatter := fmt.Sprintf(SuccessColor, message)
		log.Println(formatter)
	}
}

// Error - log error messages
func Error(message interface{}) {
	SetLogFormat()
	if isErrorLog(GetLogLevel()) {
		formatter := fmt.Sprintf(ErrorColor, message)
		log.Fatal(formatter)
	}
}