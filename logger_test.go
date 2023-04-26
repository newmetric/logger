package logger_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/newmetric/logger"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	logPrint func(msg string, args ...interface{})
	msg      string
	expected string
}

var outBuffer = &bytes.Buffer{}

func runTest(t *testing.T, tc []testCase) {
	for _, c := range tc {
		outBuffer.Reset()
		c.logPrint(c.msg)
		assert.Equal(t, c.expected, outBuffer.String())
	}
}

func TestLogger(t *testing.T) {
	assert.NoError(t, os.Setenv("TEST1_LOG_LEVEL", "info"))
	log1 := logger.SetupZeroLogger("test1", outBuffer)

	assert.NoError(t, os.Setenv("TEST2_LOG_LEVEL", "debug"))
	log2 := logger.SetupZeroLogger("test2", outBuffer)

	assert.NoError(t, os.Setenv("TEST3_LOG_LEVEL", "info"))
	optLogger := logger.SetupZeroLogger(
		"test3",
		outBuffer,
		// overwrite log level
		func(l *zerolog.Logger) *zerolog.Logger {
			newLogger := l.Level(zerolog.DebugLevel)
			return &newLogger
		})

	// For testing purpose, we set the time format to DateOnly
	zerolog.TimeFieldFormat = "2006-01-02"
	now := time.Now().Format("2006-01-02")

	testCases := []testCase{
		{log1.Debug, "debug msg", ""},
		{log1.Info, "info msg", fmt.Sprintf("{\"level\":\"info\",\"module\":\"test1\",\"time\":\"%s\",\"message\":\"info msg\"}\n", now)},
		{log2.Debug, "debug msg", fmt.Sprintf("{\"level\":\"debug\",\"module\":\"test2\",\"time\":\"%s\",\"message\":\"debug msg\"}\n", now)},
		{log2.Info, "info msg", fmt.Sprintf("{\"level\":\"info\",\"module\":\"test2\",\"time\":\"%s\",\"message\":\"info msg\"}\n", now)},
		{optLogger.Debug, "debug msg", fmt.Sprintf("{\"level\":\"debug\",\"module\":\"test3\",\"time\":\"%s\",\"message\":\"debug msg\"}\n", now)},
	}

	runTest(t, testCases)

}
