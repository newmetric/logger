package noop

type NoOpLogger struct{}

func (n *NoOpLogger) Level(string) error { return nil }

func (n *NoOpLogger) Debug(msg string, args ...interface{}) {}
func (n *NoOpLogger) Info(msg string, args ...interface{})  {}
func (n *NoOpLogger) Warn(msg string, args ...interface{})  {}
func (n *NoOpLogger) Error(msg string, args ...interface{}) {}
func (n *NoOpLogger) Trace(msg string, args ...interface{}) {}
