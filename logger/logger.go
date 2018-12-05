package logger

import (
	"log"
)

// Logger is a custom logger using the verbose flag
type Logger struct {
	verbose bool
}

var l *Logger = nil

// SetVerbose configures the verbose field
func (l *Logger) SetVerbose(verbose bool) {
	l.verbose = verbose
}

// Log displays the message to stdout if `verbose` is true
func (l *Logger) Log(v ...interface{}) {
	if l.verbose {
		log.Println(v...)
	}
}

// NewLogger instantiates a new logger
func NewLogger(verbose bool) *Logger {
	l = &Logger{}
	l.SetVerbose(verbose)
	return l
}

// GetLogger returns the singleton logger
func GetLogger() *Logger {
	return l
}
