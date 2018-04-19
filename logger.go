package ldapserver

/*
The logger interfaces is used to separate different severity of logging
within the code. If no logger is passed a default Go logger printing to
STDOUT will be used.

It is possible to use a standard logger and make it implement the interface
by passing it to the SetStdLogger() function.

If a default logger is used, all debug messages will be discarded. To make
debug logs visible, make sure the logger passed to SetLogger() implements
Debug() and Debugf().
*/

import (
	"log"
	"os"
)

var logger Logger

// Logger is an interface that implements Info(f), Warn(f) and Debug(f).
type Logger interface {
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
}

// StdLogger is a standard logger used if none is setup.
type StdLogger struct {
	*log.Logger
}

func init() {
	logger = &StdLogger{
		log.New(os.Stdout, "", log.LstdFlags),
	}
}

// SetStdLogger will use a standard Go logger and use as the standard logger.
func SetStdLogger(l *log.Logger) {
	logger = &StdLogger{l}
}

// SetLogger will set the package logger.
func SetLogger(l Logger) {
	logger = l
}

// Info is used to satisfy the logger interface with the standard logger.
func (s *StdLogger) Info(v ...interface{}) {
	s.Print(v...)
}

// Infof is used to satisfy the logger interface with the standard logger.
func (s *StdLogger) Infof(format string, v ...interface{}) {
	s.Printf(format, v...)
}

// Warn is used to satisfy the logger interface with the standard logger.
func (s *StdLogger) Warn(v ...interface{}) {
	s.Print(v...)
}

// Warnf is used to satisfy the logger interface with the standard logger.
func (s *StdLogger) Warnf(format string, v ...interface{}) {
	s.Printf(format, v...)
}

// Debug is used to satisfy the logger interface with the standard logger.
func (s *StdLogger) Debug(v ...interface{}) {
	// Debug is not ment for STDOUT.
}

// Debugf is used to satisfy the logger interface with the standard logger.
func (s *StdLogger) Debugf(format string, v ...interface{}) {
	// Debug is not ment for STDOUT.
}
