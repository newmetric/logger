package tests

import (
	"bytes"
	"os"
	"testing"

	"github.com/newmetric/logger"
	"github.com/stretchr/testify/assert"
)

type (
	testCase struct {
		logPrint func(msg string, args ...interface{})
		msg      string
		expected string
	}
)

var outBuffer = &bytes.Buffer{}

func runTests(t *testing.T, tc []testCase) {
	for _, c := range tc {
		outBuffer.Reset()
		c.logPrint(c.msg)
		assert.Equal(t, c.expected, outBuffer.String())
	}
}

func RunTests(t *testing.T, createLogger func(bufWriter *bytes.Buffer, moduleName string) logger.Logger) {
	assert.NoError(t, os.Setenv("TEST1_LOG_LEVEL", "info"))
	log1 := createLogger(outBuffer, "test1")
	log1_1 := log1.With("key1", "value1")

	testCases := []testCase{
		{log1.Debug, "debug msg", ""},
		{log1.Info, "info msg", "{\"level\":\"info\",\"module\":\"test1\",\"message\":\"info msg\"}\n"},
		{log1_1.Debug, "debug msg", ""},
		{log1_1.Info, "info msg", "{\"level\":\"info\",\"module\":\"test1\",\"key1\":\"value1\",\"message\":\"info msg\"}\n"},
	}
	runTests(t, testCases)

	assert.NoError(t, os.Setenv("TEST2_LOG_LEVEL", "debug"))
	log2 := createLogger(outBuffer, "test2")

	testCases = []testCase{
		{log2.Debug, "debug msg", "{\"level\":\"debug\",\"module\":\"test2\",\"message\":\"debug msg\"}\n"},
		{log2.Info, "info msg", "{\"level\":\"info\",\"module\":\"test2\",\"message\":\"info msg\"}\n"},
	}
	runTests(t, testCases)
}
