package noop

import "github.com/newmetric/logger/types"

type NoOpLogger struct{}

var (
	_ types.Logger = (*NoOpLogger)(nil)
)

func (n *NoOpLogger) With(args ...interface{}) types.Logger { return n }

func (n *NoOpLogger) Level(string) error { return nil }

func (n *NoOpLogger) Debug(msg string, args ...interface{}) {}
func (n *NoOpLogger) Info(msg string, args ...interface{})  {}
func (n *NoOpLogger) Warn(msg string, args ...interface{})  {}
func (n *NoOpLogger) Error(msg string, args ...interface{}) {}
func (n *NoOpLogger) Trace(msg string, args ...interface{}) {}
