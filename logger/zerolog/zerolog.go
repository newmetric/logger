package zerolog

import (
	"fmt"
	"io"
	"runtime/debug"

	"github.com/newmetric/logger/types"
	"github.com/newmetric/logger/utils"
	"github.com/rs/zerolog"
)

type ZeroLogger struct {
	w io.Writer
	*zerolog.Logger
	module string

	disableStackTrace bool
}

var (
	_ types.Logger                  = (*ZeroLogger)(nil)
	_ types.DisableStackTraceOption = (*ZeroLogger)(nil)
)

type Opts = func(*zerolog.Logger) *zerolog.Logger

func New(w io.Writer, module string, opts ...Opts) *ZeroLogger {
	logLevel := utils.GetLogLevelFromEnv(module)
	level, err := types.ParseLevel(logLevel)
	if err != nil {
		panic(err)
	}

	// create new instance
	logger := zerolog.New(w).
		Level(zerolog.Level(level)).
		With().
		Str("module", module).
		Logger()

	l := &logger

	for _, opt := range opts {
		l = opt(l)
	}

	return &ZeroLogger{
		w:      w,
		Logger: l,
		module: module,

		disableStackTrace: false,
	}
}

func (z *ZeroLogger) DisableStackTrace() {
	z.disableStackTrace = true
}

func (z *ZeroLogger) With(args ...interface{}) types.Logger {
	newLogger := z.Logger.With().Fields(args).Logger()

	return &ZeroLogger{
		Logger: &newLogger,
	}
}

func (z *ZeroLogger) SetLevel(level types.Level) error {
	newLogger := z.Logger.Level(zerolog.Level(level))
	z.Logger = &newLogger

	return nil
}

func (z *ZeroLogger) GetLevel() types.Level {
	return types.Level(z.Logger.GetLevel())
}

func (z *ZeroLogger) Debug(msg string, args ...interface{}) {
	z.Logger.Debug().Fields(args).Msg(msg)
}

func (z *ZeroLogger) Info(msg string, args ...interface{}) {
	z.Logger.Info().Fields(args).Msg(msg)
}

func (z *ZeroLogger) Warn(msg string, args ...interface{}) {
	z.Logger.Warn().Fields(args).Msg(msg)
}

func (z *ZeroLogger) Error(msg string, args ...interface{}) {
	var stackArg []interface{} = make([]interface{}, 0)
	if !z.disableStackTrace {
		stackArg = []interface{}{"stack-trace", debug.Stack()}
	}

	if len(args)%2 == 1 { // if odd(It shouldn't be, but that's how external library writes it.)
		for _, arg := range args {
			switch v := arg.(type) {
			case error:
				msg += fmt.Sprintf("| %s", v.Error())
			case string:
				msg += fmt.Sprintf("| %s", v)
			}
		}
		z.Logger.Error().Fields(stackArg).Msg(msg)
	} else {
		z.Logger.Error().Fields(args).Fields(stackArg).Msg(msg)
	}
}

func (z *ZeroLogger) Fatal(msg string, args ...interface{}) {
	var stackArg []interface{} = make([]interface{}, 0)
	if !z.disableStackTrace {
		stackArg = []interface{}{"stack-trace", debug.Stack()}
	}

	if len(args)%2 == 1 { // if odd(It shouldn't be, but that's how external library writes it.)
		for _, arg := range args {
			switch v := arg.(type) {
			case error:
				msg += fmt.Sprintf("| %s", v.Error())
			case string:
				msg += fmt.Sprintf("| %s", v)
			}
		}
		z.Logger.Fatal().Fields(stackArg).Msg(msg)
	} else {
		z.Logger.Fatal().Fields(args).Fields(stackArg).Msg(msg)
	}
}

func (z *ZeroLogger) Trace(msg string, args ...interface{}) {
	z.Logger.Trace().Fields(args).Msg(msg)
}
