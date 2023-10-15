package logwrapper

import (
	log "github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	*log.Entry
}

func NewLogrusLogger(name string) Logger {
	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{})
	return &LogrusLogger{logger.WithField("logger_name", name)}
}

func (l *LogrusLogger) Debug(msg string, keysAndValues ...interface{}) {
	l.Entry.Debugf(msg, keysAndValues...)
}

func (l *LogrusLogger) Info(msg string, keysAndValues ...interface{}) {
	l.Entry.Infof(msg, keysAndValues...)
}

func (l *LogrusLogger) Error(msg string, keysAndValues ...interface{}) {
	l.Entry.Errorf(msg, keysAndValues...)
}

func (l *LogrusLogger) Fatal(msg string, keysAndValues ...interface{}) {
	l.Entry.Fatalf(msg, keysAndValues...)
}
