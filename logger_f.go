package golog

import (
	"fmt"
	"log"
	"runtime"

	"github.com/fatih/color"
)

// Debugf - log debugging messages in same line
func Debugf(message interface{}) {
	debug := color.New(color.FgCyan).SprintfFunc()
	var fileInfo string = "<nil>"

	_, file, lineNum, ok := runtime.Caller(1)

	if ok {
		fileInfo = fmt.Sprintf("%s::%d", file, lineNum)
	}

	if isDebugLog(GetLogLevel()) {
		log.Printf(debug("[DEBUG] %s %s", fileInfo, message))
	}
}

// Infof - log informative messages in same line
func Infof(message interface{}) {
	info := color.New(color.FgBlue).SprintfFunc()

	if isInfoLog(GetLogLevel()) {
		log.Printf(info("[INFO] %v", message))
	}
}

// Warnf - log warning messages in same line
func Warnf(message interface{}) {
	warn := color.New(color.FgBlack).Add(color.BgYellow).SprintfFunc()

	if isWarnLog(GetLogLevel()) {
		log.Printf(warn("[WARNING] %v", message))
	}
}

// Successf - log success messages in same line
func Successf(message interface{}) {
	success := color.New(color.FgBlack).Add(color.BgGreen).SprintfFunc()

	if isSuccessLog(GetLogLevel()) {
		log.Printf(success("[SUCCESS] %v", message))
	}
}

// Errorf - log error messages in same line
func Errorf(message interface{}) {
	error := color.New(color.FgRed).Add(color.Bold).SprintfFunc()

	if isErrorLog(GetLogLevel()) {
		log.Fatalf(error("[ERROR] %v", message))
	}
}

// Printf - print logs without any constraints
func Printf(message interface{}) {
	log.Printf(fmt.Sprintf("%v", message))
}