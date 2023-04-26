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

// logger

func SetupZeroLogger(module string, w io.Writer, opts ...zerolog.Opts) types.Logger {
	logger := zerolog.New(w, module, opts...)
	LoggerMap[module] = logger

	return logger
}

func SetupNoOpLogger() types.Logger {
	return &noop.NoOpLogger{}
}
