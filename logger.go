package logger

import (
	"fmt"
	"io"

	"github.com/newmetric/logger/logger/noop"
	"github.com/newmetric/logger/logger/zerolog"
	"github.com/newmetric/logger/types"
)

// re-exports
type (
	// instance type
	Logger = types.Logger

	// option
	DisableStackTraceOption = types.DisableStackTraceOption
)

const (
	Debug = types.DebugLevel
	Info  = types.InfoLevel
	Warn  = types.WarnLevel
	Error = types.ErrorLevel
	Fatal = types.FatalLevel

	Disabled = types.Disabled

	Trace = types.TraceLevel
)

var (
	ParseLevel = types.ParseLevel
)

// ==========

var LoggerMap map[string]types.Logger = make(map[string]types.Logger)

func ChangeLevel(module string, level types.Level) error {
	logger, ok := LoggerMap[module]
	if !ok {
		return fmt.Errorf("logger: module %s not found", module)
	}

	return logger.SetLevel(level)
}

// logger instance

func SetupZeroLogger(module string, w io.Writer, opts ...zerolog.Opts) types.Logger {
	logger := zerolog.New(w, module, opts...)
	LoggerMap[module] = logger

	return logger
}

func SetupNoOpLogger() types.Logger {
	return &noop.NoOpLogger{}
}

// logger option helper

func DisableStackTrace(logger Logger) error {
	if logger, ok := logger.(DisableStackTraceOption); ok {
		logger.DisableStackTrace()
		return nil
	}

	return fmt.Errorf("logger: logger %T does not support disable stack trace option", logger)
}
