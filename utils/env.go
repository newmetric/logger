package utils

import (
	"fmt"
	"os"
	"strings"
)

const (
	defaultLogLevel = "info"
)

func GetLogLevelFromEnv(module string) string {
	logLevel := os.Getenv(fmt.Sprintf("%s_LOG_LEVEL", strings.ToUpper(module)))

	if logLevel == "" {
		logLevel = defaultLogLevel
	}

	return logLevel
}
