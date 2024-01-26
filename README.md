# newmetric/logger

An implementation that connects multiple logger implementations with a common interface

## Logger common interface

```go
type Logger interface {
	// SetLevel sets the logger level
	SetLevel(level Level) error

	// GetLevel returns the current logger level
	GetLevel() Level

	// ReplaceOutputWriter replaces the output writer of the current logger
	ReplaceOutputWriter(w io.Writer)

	// With returns a new Logger with keyvals prepended to those passed to calls to
	With(args ...interface{}) Logger

	// Debug prints debug log
	Debug(msg string, args ...interface{})

	// Info prints info log
	Info(msg string, args ...interface{})

	// Warn prints warn log
	Warn(msg string, args ...interface{})

	// Error prints error log
	Error(msg string, args ...interface{})

	// Fatal prints fatal log
	Fatal(msg string, args ...interface{})

	// Trace prints trace log
	Trace(msg string, args ...interface{})
}
```

## Logger extension

```go
type DisableStackTraceOption interface {
	DisableStackTrace()
}

type FormatLogger interface {
	types.Logger

	// sort by log level
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
	Tracef(format string, args ...any)
}
```
