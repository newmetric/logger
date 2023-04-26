package logger

import (
	"fmt"
	"io"

	"github.com/newmetric/logger/noop"
	"github.com/newmetric/logger/zerolog"
)

var LoggerMap map[string]Logger = make(map[string]Logger)

var (
	_ Logger = (*zerolog.ZeroLogger)(nil)
	_ Logger = (*noop.NoOpLogger)(nil)
)

func ChangeLevel(module string, level string) error {
	logger, ok := LoggerMap[module]
	if !ok {
		return fmt.Errorf("logger: module %s not found", module)
	}

	return logger.Level(level)
}

type Logger interface {
	Level(level string) error

	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Trace(msg string, args ...interface{})
}

func SetupZeroLogger(module string, w io.Writer, opts ...zerolog.Opts) Logger {
	logger := zerolog.New(w, module, opts...)
	LoggerMap[module] = logger

	return logger
}

func SetupNoOpLogger() Logger {
	return &noop.NoOpLogger{}
}
