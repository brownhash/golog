package golog

import (
	"fmt"
	"log"
	"runtime"

	"github.com/fatih/color"
)

// Debug - log debugging messages
func Debug(message interface{}) {
	debug := color.New(color.FgCyan).SprintfFunc()
	var fileInfo string = "<nil>"

	_, file, lineNum, ok := runtime.Caller(1)

	if ok {
		fileInfo = fmt.Sprintf("%s::%d", file, lineNum)
	}

	if isDebugLog(GetLogLevel()) {
		log.Println(debug("[DEBUG] %s %s", fileInfo, message))
	}
}

// Info - log informative messages
func Info(message interface{}) {
	info := color.New(color.FgBlue).SprintfFunc()

	if isInfoLog(GetLogLevel()) {
		log.Println(info("[INFO] %v", message))
	}
}

// Warn - log warning messages
func Warn(message interface{}) {
	warn := color.New(color.FgBlack).Add(color.BgYellow).SprintfFunc()

	if isWarnLog(GetLogLevel()) {
		log.Println(warn("[WARNING] %v", message))
	}
}

// Success - log success messages
func Success(message interface{}) {
	success := color.New(color.FgBlack).Add(color.BgGreen).SprintfFunc()

	if isSuccessLog(GetLogLevel()) {
		log.Println(success("[SUCCESS] %v", message))
	}
}

// Error - log error messages
func Error(message interface{}) {
	error := color.New(color.FgRed).Add(color.Bold).SprintfFunc()

	if isErrorLog(GetLogLevel()) {
		log.Fatal(error("[ERROR] %v", message))
	}
}

// Println - print logs without any constraints
func Println(message interface{}) {
	log.Println(fmt.Sprintf("%v", message))
}