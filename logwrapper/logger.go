package logwrapper

// LoggerType constants represent the type of logger to be created.
const (
	LogrusLoggerType = "logrus"
	ZapLoggerType    = "zap"
	CharmLoggerType  = "charm"
	StdLoggerType    = "std"
)

// Logger is a generic interface for logging with different loggers.
type Logger interface {
	Debug(msg string, keysAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
	Fatal(msg string, keysAndValues ...interface{})
}

// NewLogger is a factory function that creates a new logger based on the given type.
func NewLogger(loggerType string, name string) Logger {
	switch loggerType {
	case LogrusLoggerType:
		return NewLogrusLogger(name)
	case ZapLoggerType:
		return NewZapLogger(name)
	case CharmLoggerType:
		return NewCharmLogger(name)
	case StdLoggerType:
		return NewStdLogger(name)
	default:
		return NewStdLogger(name)
	}
}
