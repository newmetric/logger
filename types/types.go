package types

import (
	"fmt"
	"io"
)

type Level int8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel

	Disabled

	TraceLevel Level = -1
)

func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	case Disabled:
		return "disabled"
	case TraceLevel:
		return "trace"
	default:
		return "unknown"
	}
}

type Logger interface {
	// SetLevel sets the logger level
	SetLevel(level Level) error

	// GetLevel returns the current logger level
	GetLevel() Level

	// ReplaceOutputWriter replaces the output writer of the current logger
	ReplaceOutputWriter(w io.Writer)

	// like fmt.Printf
	Printf(format string, args ...interface{})

	// With returns a new Logger with keyvals prepended to those passed to calls to
	With(args ...interface{}) Logger

	// Debug prints debug log
	Debug(msg string, args ...interface{})

	// Info prints info log
	Info(msg string, args ...interface{})

	// Warn prints warn log
	Warn(msg string, args ...interface{})

	// Error prints error log
	Error(msg string, args ...interface{})

	// Fatal prints fatal log
	Fatal(msg string, args ...interface{})

	// Trace prints trace log
	Trace(msg string, args ...interface{})
}

type DisableStackTraceOption interface {
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
