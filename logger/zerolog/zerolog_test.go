package zerolog_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/newmetric/logger/logger/tests"
	"github.com/newmetric/logger/logger/zerolog"
	"github.com/newmetric/logger/types"
	"github.com/stretchr/testify/assert"
)

func TestZeroLog(t *testing.T) {
	tests.RunTests(
		t,
		func(bufWriter *bytes.Buffer, moduleName string) types.Logger {
			return zerolog.New(bufWriter, moduleName)
		},
	)
}

func TestZeroLogStackTrace(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := zerolog.New(buf, "test")

	logger.Error("error msg")

	contains := func(s string) {
		assert.True(t, strings.Contains(buf.String(), s))
	}

	contains("logger/logger/zerolog.(*ZeroLogger).Error(")
	contains("runtime/debug.Stack()")
}
