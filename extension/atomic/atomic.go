package atomic

import (
	"encoding/json"
	"io"
	"net/http"
	"sync/atomic"

	"github.com/newmetric/logger/types"
)

type AtomicLogger interface {
	// atomic change of log level
	SetLevel(level types.Level) error
	GetLevel() types.Level

	HttpHandler() http.HandlerFunc

	types.Logger
}

type logger struct {
	level *atomic.Int32

	types.Logger
}

var (
	_ types.Logger = (*logger)(nil)
	_ AtomicLogger = (*logger)(nil)
)

func New(l types.Logger) *logger {
	level := &atomic.Int32{}
	level.Store(int32(l.GetLevel()))

	l.SetLevel(types.TraceLevel)

	return &logger{
		level:  level,
		Logger: l,
	}
}

func (logger *logger) HttpHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		body, _ := io.ReadAll(r.Body)

		m := make(map[string]string)
		err := json.Unmarshal(body, &m)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		level, ok := m["level"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// parse level
		l, err := types.ParseLevel(level)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// set level
		err = logger.SetLevel(l)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}
}

func (l *logger) With(args ...interface{}) types.Logger {
	newLogger := l.Logger.With(args...)

	return &logger{
		level:  l.level,
		Logger: newLogger,
	}
}

func (l *logger) SetLevel(level types.Level) error {
	l.level.Store(int32(level))
	return nil
}
func (l *logger) GetLevel() types.Level {
	return types.Level(l.level.Load())
}

func (l *logger) Debug(msg string, args ...interface{}) {
	if l.GetLevel() <= types.DebugLevel {
		l.Logger.Debug(msg, args...)
	}
}

func (l *logger) Info(msg string, args ...interface{}) {
	if l.GetLevel() <= types.InfoLevel {
		l.Logger.Info(msg, args...)
	}
}

func (l *logger) Warn(msg string, args ...interface{}) {
	if l.GetLevel() <= types.WarnLevel {
		l.Logger.Warn(msg, args...)
	}
}

func (l *logger) Error(msg string, args ...interface{}) {
	if l.GetLevel() <= types.ErrorLevel {
		l.Logger.Error(msg, args...)
	}
}

func (l *logger) Fatal(msg string, args ...interface{}) {
	if l.GetLevel() <= types.FatalLevel {
		l.Logger.Fatal(msg, args...)
	}
}

func (l *logger) Trace(msg string, args ...interface{}) {
	if l.GetLevel() <= types.TraceLevel {
		l.Logger.Trace(msg, args...)
	}
}
