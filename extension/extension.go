package extension

import (
	"github.com/newmetric/logger/extension/fmt"
	"github.com/newmetric/logger/types"
)

func ApplyFmtExtensino(l types.Logger) fmt.FormatLogger {
	return fmt.New(l)
}
