package noop

import "github.com/newmetric/logger/types"

type NoOpLogger struct{}

var (
	_ types.Logger = (*NoOpLogger)(nil)
)

func (n *NoOpLogger) With(args ...interface{}) types.Logger { return n }

func (n *NoOpLogger) SetLevel(types.Level) error { return nil }
func (n *NoOpLogger) GetLevel() types.Level      { return types.Disabled }

func (n *NoOpLogger) Debug(msg string, args ...interface{}) {}
func (n *NoOpLogger) Info(msg string, args ...interface{})  {}
func (n *NoOpLogger) Warn(msg string, args ...interface{})  {}
func (n *NoOpLogger) Error(msg string, args ...interface{}) {}
func (n *NoOpLogger) Fatal(msg string, args ...interface{}) {}
func (n *NoOpLogger) Trace(msg string, args ...interface{}) {}
