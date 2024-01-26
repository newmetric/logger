package atomic_test

import (
	"io"
	"sync"
	std_atomic "sync/atomic"
	"testing"

	"github.com/newmetric/logger/extension/atomic"
	"github.com/newmetric/logger/logger/zerolog"
	"github.com/newmetric/logger/types"
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
