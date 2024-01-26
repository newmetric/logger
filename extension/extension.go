package extension

import (
	"github.com/newmetric/logger/extension/atomic"
	"github.com/newmetric/logger/extension/fmt"
	"github.com/newmetric/logger/types"
)

func ApplyAtomicExtension(l types.Logger) atomic.AtomicLogger {
	return atomic.New(l)
}

func ApplyFmtExtension(l types.Logger) fmt.FormatLogger {
	return fmt.New(l)
}
