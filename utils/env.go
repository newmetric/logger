package utils

import (
	"fmt"
	"os"
	"strings"
)

const (
	defaultLogLevel = "info"
)

func GetLogLevelFromEnvPerModule(module string) string {
	// replace dash to underscore
	module = strings.ReplaceAll(module, "-", "_")

	logLevel := os.Getenv(fmt.Sprintf("%s_LOG_LEVEL", strings.ToUpper(module)))

	if logLevel == "" {
		logLevel = defaultLogLevel
	}

	return logLevel
}

func GetLogLevelFromEnv() string {
	logLevel := os.Getenv("LOG_LEVEL")

	if logLevel == "" {
		logLevel = defaultLogLevel
	}

	return logLevel
}
