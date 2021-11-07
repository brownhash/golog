package golog

import (
	"testing"
)

func TestLoggers(t *testing.T) {
	// Set log level to lowest
	SetLogLevel("DEBUG")
	// Ste logging format
	SetLogFormat()

	// Debug log
    Debug("Debug log!")
	Debugf("Debugf log!")

	// Info log
	Info("Info log!")
	Infof("Infof log!")

	// Warn log
	Warn("Warning log!")
	Warnf("Warningf log!")

	// Success log
	Success("Success log!")
	Successf("Successf log!")

	// Print log
	Println("Print log!")
	Printf("Printf log!")
	// do not call Error()
}
