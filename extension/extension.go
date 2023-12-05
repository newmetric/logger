package extension

import (
	"github.com/newmetric/logger/extension/fmt"
	"github.com/newmetric/logger/types"
)

func ApplyFmtExtension(l types.Logger) fmt.FormatLogger {
	return fmt.New(l)
}
