package logwrapper

import (
	"log"
	"os"
)

type StdLogger struct {
	name string
}

func NewStdLogger(name string) Logger {
	return &StdLogger{name: name}
}

func (l *StdLogger) Debug(msg string, keysAndValues ...interface{}) {
	l.log("DEBUG", msg, keysAndValues...)
}

func (l *StdLogger) Info(msg string, keysAndValues ...interface{}) {
	l.log("INFO", msg, keysAndValues...)
}

func (l *StdLogger) Error(msg string, keysAndValues ...interface{}) {
	l.log("ERROR", msg, keysAndValues...)
}

func (l *StdLogger) Fatal(msg string, keysAndValues ...interface{}) {
	l.log("FATAL", msg, keysAndValues...)
	os.Exit(1)
}

func (l *StdLogger) log(level string, msg string, keysAndValues ...interface{}) {
	log.Printf("%s [%s] %s %v", level, l.name, msg, keysAndValues)
}
