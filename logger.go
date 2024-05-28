package logger

import (
	"fmt"
	"io"

	"github.com/newmetric/logger/logger/noop"
	"github.com/newmetric/logger/logger/zerolog"
	"github.com/newmetric/logger/types"
	"github.com/newmetric/logger/utils"
)

// re-exports
type (
	Logger                  = types.Logger
	DisableStackTraceOption = types.DisableStackTraceOption

	Level = types.Level
)

var (
	// ParseLevel parses a level string into a logger Level value.
	ParseLevel = types.ParseLevel

	// LoggerMap is a map managing levels for individual modules
	LoggerMap map[string]types.Logger = make(map[string]types.Logger)
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

// ChangeLevel changes the level of a logger.
func ChangeLevel(module string, level types.Level) error {
	logger, ok := LoggerMap[module]
	if !ok {
		return fmt.Errorf("logger: module %s not found", module)
	}

	return logger.SetLevel(level)
}

// SetupZeroLogger returns a new logger instance.
func SetupZeroLogger(module string, w io.Writer, opts ...zerolog.Opts) types.Logger {
	logger := zerolog.New(w, module, opts...)
	logLevel := utils.GetLogLevelFromEnv()
	level, err := ParseLevel(logLevel)
	if err != nil {
		panic(err) // unreachable
	}

	logger.SetLevel(level)
	LoggerMap[module] = logger

	return logger
}

// SetupNoOpLogger returns a no-op logger.
func SetupNoOpLogger() types.Logger {
	return &noop.NoOpLogger{}
}

// DisableStackTrace disables stack trace for logger.
func DisableStackTrace(logger Logger) error {
	if logger, ok := logger.(DisableStackTraceOption); ok {
		logger.DisableStackTrace()
		return nil
	}

	return fmt.Errorf("logger: logger %T does not support disable stack trace option", logger)
}
