package logger

import (
	"io"

	"github.com/newmetric/logger/zerolog"
)

var LoggerMap map[string]Logger = make(map[string]Logger)

var (
	_ Logger = (*zerolog.ZeroLogger)(nil)
)

type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Trace(msg string, args ...interface{})
}

func SetupZeroLog(module string, w io.Writer, opts ...zerolog.Opts) Logger {
	logger := zerolog.New(w, module, opts...)
	LoggerMap[module] = logger

	return logger
}
