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
	Logger = types.Logger
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
