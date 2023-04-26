package zerolog_test

import (
	"bytes"
	"testing"

	"github.com/newmetric/logger/logger/tests"
	"github.com/newmetric/logger/logger/zerolog"
	"github.com/newmetric/logger/types"
)

func TestZeroLog(t *testing.T) {
	tests.RunTests(
		t,
		func(bufWriter *bytes.Buffer, moduleName string) types.Logger {
			return zerolog.New(bufWriter, moduleName)
		},
	)
}
