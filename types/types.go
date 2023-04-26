package types

type Level = int8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel

	Disabled

	TraceLevel Level = -1
)

type Logger interface {
	SetLevel(level Level) error
	GetLevel() Level

	With(args ...interface{}) Logger

	// sort by log level
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
	Trace(msg string, args ...interface{})
}
