package types

type Logger interface {
	Level(level string) error
	With(args ...interface{}) Logger

	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Trace(msg string, args ...interface{})
}
