package types

import "fmt"

type Level = int8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel

	Disabled

	TraceLevel Level = -1
)

type Logger interface {
	SetLevel(level Level) error
	GetLevel() Level

	With(args ...interface{}) Logger

	// sort by log level
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
	Trace(msg string, args ...interface{})
}

type StackTraceOption interface {
	DisableStackTrace()
}

func ParseLevel(level string) (Level, error) {
	switch level {
	case "debug":
		return DebugLevel, nil
	case "info":
		return InfoLevel, nil
	case "warn":
		return WarnLevel, nil
	case "error":
		return ErrorLevel, nil
	case "fatal":
		return FatalLevel, nil

	case "disabled":
		return Disabled, nil

	case "trace":
		return TraceLevel, nil

	default:
		return Disabled, fmt.Errorf("logger: invalid log level %s", level)
	}
}
