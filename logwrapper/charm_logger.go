package logwrapper

import (
	"github.com/charmbracelet/log"
	"os"
)

type CharmLogger struct {
	*log.Logger
}

func NewCharmLogger(name string) Logger {
	logger := log.New(os.Stderr)

	return &CharmLogger{
		Logger: logger.With("logger_name", name),
	}
}

func (c *CharmLogger) Debug(msg string, keysAndValues ...interface{}) {
	log.Debug(msg, keysAndValues...)
}

func (c *CharmLogger) Info(msg string, keysAndValues ...interface{}) {
	log.Info(msg, keysAndValues...)
}

func (c *CharmLogger) Error(msg string, keysAndValues ...interface{}) {
	log.Error(msg, keysAndValues...)
}

func (c *CharmLogger) Fatal(msg string, keysAndValues ...interface{}) {
	log.Fatal(msg, keysAndValues...)
}
