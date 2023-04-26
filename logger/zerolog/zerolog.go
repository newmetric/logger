package zerolog

import (
	"io"

	"github.com/newmetric/logger/types"
	"github.com/newmetric/logger/utils"
	"github.com/rs/zerolog"
)

type ZeroLogger struct {
	*zerolog.Logger
	module string
}

var (
	_ types.Logger = (*ZeroLogger)(nil)
)

type Opts = func(*zerolog.Logger) *zerolog.Logger

func New(w io.Writer, module string, opts ...Opts) *ZeroLogger {
	logLevel := utils.GetLogLevelFromEnv(module)
	level, err := types.ParseLevel(logLevel)
	if err != nil {
		panic(err)
	}

	// create new instance
	logger := zerolog.New(w).Level(zerolog.Level(level)).With().
		Str("module", module).
		Logger()

	l := &logger

	for _, opt := range opts {
		l = opt(l)
	}

	zlogger := &ZeroLogger{
		Logger: l,
		module: module,
	}

	return zlogger
}

func (z *ZeroLogger) With(args ...interface{}) types.Logger {
	newLogger := z.Logger.With().Fields(args).Logger()
	z.Logger = &newLogger
	return z
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
	z.Logger.Error().Fields(args).Msg(msg)
}

func (z *ZeroLogger) Fatal(msg string, args ...interface{}) {
	z.Logger.Fatal().Fields(args).Msg(msg)
}

func (z *ZeroLogger) Trace(msg string, args ...interface{}) {
	z.Logger.Trace().Fields(args).Msg(msg)
}
