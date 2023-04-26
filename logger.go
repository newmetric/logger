package logger

import (
	"fmt"
	"io"

	"github.com/newmetric/logger/noop"
	"github.com/newmetric/logger/types"
	"github.com/newmetric/logger/zerolog"
)

var LoggerMap map[string]types.Logger = make(map[string]types.Logger)

func ChangeLevel(module string, level string) error {
	logger, ok := LoggerMap[module]
	if !ok {
		return fmt.Errorf("logger: module %s not found", module)
	}

	return logger.Level(level)
}

func SetupZeroLogger(module string, w io.Writer, opts ...zerolog.Opts) types.Logger {
	logger := zerolog.New(w, module, opts...)
	LoggerMap[module] = logger

	return logger
}

func SetupNoOpLogger() types.Logger {
	return &noop.NoOpLogger{}
}
