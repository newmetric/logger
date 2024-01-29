package atomic_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	std_atomic "sync/atomic"
	"testing"

	"github.com/newmetric/logger/extension/atomic"
	"github.com/newmetric/logger/logger/noop"
	"github.com/newmetric/logger/logger/zerolog"
	"github.com/newmetric/logger/types"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/assert"
)

type arena struct {
	buf []byte
	n   *std_atomic.Int32
}

func newArena(size int) *arena {
	n := std_atomic.Int32{}
	n.Store(1)

	return &arena{
		buf: make([]byte, size),
		n:   &n,
	}
}

var _ (io.Writer) = (*arena)(nil)

func (a *arena) Write(p []byte) (n int, err error) {
	next := a.n.Add(int32(len(p)))
	if int(next) > len(a.buf) {
		return 0, io.ErrShortBuffer
	}

	offset := next - int32(len(p))
	copy(a.buf[offset:next], p)
	return len(p), nil
}

func TestAtomicLoggerRace(t *testing.T) {
	wg := &sync.WaitGroup{}

	run := func(cb func()) {
		for i := 0; i < 1; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				cb()
			}()
		}
	}

	buf := newArena(1024 * 5)
	logger := zerolog.New(buf, "test")
	atomicLogger := atomic.New(logger)

	run(func() { atomicLogger.Debug("1") })
	run(func() { atomicLogger.Info("2") })
	run(func() { atomicLogger.Warn("3") })
	run(func() { atomicLogger.Error("4") })
	run(func() { atomicLogger.Trace("6") })
	run(func() { atomicLogger.SetLevel(types.TraceLevel) })

	wg.Wait()

	assert.NotEqual(t, buf.n.Load(), 0, "No logs were written")
}

func TestCheckLevel(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	logger := zerolog.New(buf, "test")
	atomicLogger := atomic.New(logger)

	levelCheck := func(level types.Level, tcs []struct {
		expected bool
		cb       func()
	}) {
		atomicLogger.SetLevel(level)

		for _, tc := range tcs {
			buf.Reset()
			tc.cb()
			assert.Equal(t,
				tc.expected,
				0 < buf.Len(),
				"Level %s", level.String(),
			)
		}
	}

	levelCheck(types.DebugLevel,
		[]struct {
			expected bool
			cb       func()
		}{
			{false, func() { atomicLogger.Trace("6") }},

			{true, func() { atomicLogger.Info("2") }},
			{true, func() { atomicLogger.Debug("1") }},
			{true, func() { atomicLogger.Warn("3") }},
			{true, func() { atomicLogger.Error("4") }},
		},
	)
	levelCheck(types.InfoLevel,
		[]struct {
			expected bool
			cb       func()
		}{
			{false, func() { atomicLogger.Debug("1") }},
			{false, func() { atomicLogger.Trace("6") }},

			{true, func() { atomicLogger.Warn("3") }},
			{true, func() { atomicLogger.Info("2") }},
			{true, func() { atomicLogger.Error("4") }},
		},
	)
	levelCheck(types.WarnLevel,
		[]struct {
			expected bool
			cb       func()
		}{
			{false, func() { atomicLogger.Debug("1") }},
			{false, func() { atomicLogger.Trace("6") }},
			{false, func() { atomicLogger.Info("2") }},

			{true, func() { atomicLogger.Error("4") }},
			{true, func() { atomicLogger.Warn("6") }},
		},
	)
	levelCheck(types.ErrorLevel,
		[]struct {
			expected bool
			cb       func()
		}{
			{false, func() { atomicLogger.Debug("1") }},
			{false, func() { atomicLogger.Trace("6") }},
			{false, func() { atomicLogger.Info("2") }},
			{false, func() { atomicLogger.Warn("3") }},

			{true, func() { atomicLogger.Error("4") }},
		},
	)
	levelCheck(types.TraceLevel,
		[]struct {
			expected bool
			cb       func()
		}{
			{true, func() { atomicLogger.Debug("1") }},
			{true, func() { atomicLogger.Trace("6") }},
			{true, func() { atomicLogger.Info("2") }},
			{true, func() { atomicLogger.Warn("3") }},
			{true, func() { atomicLogger.Error("4") }},
		},
	)
}

func TestAtomicLoggerHttpHandler(t *testing.T) {
	port, err := freeport.GetFreePort()
	if err != nil {
		panic(err)
	}

	addr := fmt.Sprintf("localhost:%d", port)
	logger := &noop.NoOpLogger{}
	atomicLogger := atomic.New(logger)
	atomicLogger.SetLevel(types.InfoLevel)

	go func() {
		err := http.ListenAndServe(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.URL.Path)
			if r.URL.Path == "/level" {
				atomicLogger.HttpHandler()(w, r)
				return
			}
		}))
		assert.NoError(t, err)
	}()

	req := func(method, path string, level types.Level) (*http.Response, error) {
		m := make(map[string]string)
		m["level"] = level.String()
		jsonBytes, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}

		url := fmt.Sprintf("http://%s%s", addr, path)
		request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
		if err != nil {
			return nil, err
		}

		return http.DefaultClient.Do(request)
	}

	// set level
	resp, err := req("POST", "/level", types.DebugLevel)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// check level
	assert.Equal(t, types.DebugLevel, atomicLogger.GetLevel())
}

func TestDerivedLogger(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	logger := zerolog.New(buf, "test")
	logger.SetLevel(types.InfoLevel)
	atomicLogger := atomic.New(logger)

	// derived logger
	derivedLogger := atomicLogger.With("key", "derived")

	// print debug log, but it should not be printed
	derivedLogger.Debug("1")
	assert.NotContains(t, buf.String(), "key")
	assert.NotContains(t, buf.String(), "derived")

	// print info log
	derivedLogger.Info("1")

	// verify that the newly added keys and value are printed
	assert.Contains(t, buf.String(), "key")
	assert.Contains(t, buf.String(), "derived")
}

func TestDerivedLoggerChangeLevel(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	logger := zerolog.New(buf, "test")
	logger.SetLevel(types.InfoLevel)
	atomicLogger := atomic.New(logger)

	// derived logger
	derivedLogger := atomicLogger.With("key", "derived")

	// lower origin logger level
	atomicLogger.SetLevel(types.DebugLevel)

	// print
	derivedLogger.Debug("1")

	// check whether the derived logger is affected by changes to the original logger level.
	assert.Contains(t, buf.String(), "key")
	assert.Contains(t, buf.String(), "derived")
}
