package extension

import (
	"fmt"

	"github.com/newmetric/logger/types"
)

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

type logger struct {
	types.Logger
}

var (
	_ types.Logger = (*logger)(nil)
	_ FormatLogger = (*logger)(nil)
)

func New(l types.Logger) *logger {
	return &logger{
		Logger: l,
	}
}

func (l *logger) Debugf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.Logger.Debug(msg)
}
func (l *logger) Infof(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.Logger.Info(msg)
}
func (l *logger) Warnf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.Logger.Warn(msg)
}
func (l *logger) Errorf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.Logger.Error(msg)
}
func (l *logger) Fatalf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.Logger.Fatal(msg)
}
func (l *logger) Tracef(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.Logger.Trace(msg)
}
