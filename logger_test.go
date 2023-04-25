package logger_test

import (
	"os"
	"testing"

	"github.com/newmetric/logger"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	assert.NoError(t, os.Setenv("TEST1_LOG_LEVEL", "info"))
	log1 := logger.SetupZeroLog("test1", os.Stdout)
	assert.NoError(t, os.Setenv("TEST2_LOG_LEVEL", "debug"))
	log2 := logger.SetupZeroLog("test2", os.Stdout)

	{
		// no-op
		log1.Debug("debug msg")
		//  {"level":"info","module":"test1","time":"...","message":"info msg"}
		log1.Info("info msg")
	}

	println()
	{
		//  {"level":"debug","module":"test2","time":"...","message":"debug msg"}
		log2.Debug("debug msg")
		//  {"level":"info","module":"test2","time":"...","message":"info msg"}
		log2.Info("info msg")
	}

	assert.NoError(t, os.Setenv("TEST3_LOG_LEVEL", "info"))
	optLogger := logger.SetupZeroLog(
		"test3",
		os.Stdout,
		// overwrite log level
		func(l *zerolog.Logger) *zerolog.Logger {
			newLogger := l.Level(zerolog.DebugLevel)
			return &newLogger
		})

	println()
	{
		// {"level":"debug","module":"test3","time":"...","message":"debug msg"}
		optLogger.Debug("debug msg")
	}

	assert.NoError(t, optLogger.Level("info"))
	println()
	{
		// no-op
		optLogger.Debug("debug msg")
	}
}
