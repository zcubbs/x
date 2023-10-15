package log

import (
	"log"
	"os"
)

// Logger is an interface that wraps the standard "log.Logger" methods.
type Logger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

// Level represents different logging levels.
type Level int

const (
	// InfoLevel represents the info logging level.
	InfoLevel Level = iota
	// ErrorLevel represents the error logging level.
	ErrorLevel
	// DebugLevel represents the debug logging level.
	DebugLevel
)

// Wrapper is a struct that wraps a logger and a log level.
type Wrapper struct {
	logger Logger
	level  Level
}

// NewLoggerWrapper creates a new LoggerWrapper with the specified logger and log level.
func NewLoggerWrapper(logger Logger, level Level) *Wrapper {
	return &Wrapper{
		logger: logger,
		level:  level,
	}
}

// SetLogLevel sets the log level of the logger.
func (lw *Wrapper) SetLogLevel(level Level) {
	lw.level = level
}

// Info logs an info message.
func (lw *Wrapper) Info(v ...interface{}) {
	if lw.level <= InfoLevel {
		lw.logger.Print(v...)
	}
}

// Infof logs a formatted info message.
func (lw *Wrapper) Infof(format string, v ...interface{}) {
	if lw.level <= InfoLevel {
		lw.logger.Printf(format, v...)
	}
}

// Error logs an error message.
func (lw *Wrapper) Error(v ...interface{}) {
	if lw.level <= ErrorLevel {
		lw.logger.Print(v...)
	}
}

// Errorf logs a formatted error message.
func (lw *Wrapper) Errorf(format string, v ...interface{}) {
	if lw.level <= ErrorLevel {
		lw.logger.Printf(format, v...)
	}
}

// Debug logs a debug message.
func (lw *Wrapper) Debug(v ...interface{}) {
	if lw.level <= DebugLevel {
		lw.logger.Print(v...)
	}
}

// Debugf logs a formatted debug message.
func (lw *Wrapper) Debugf(format string, v ...interface{}) {
	if lw.level <= DebugLevel {
		lw.logger.Printf(format, v...)
	}
}

// NewStandardLogger creates a new logger wrapper that uses the standard "log" package logger.
func NewStandardLogger(level Level) *Wrapper {
	return NewLoggerWrapper(log.New(os.Stdout, "", log.LstdFlags), level)
}
