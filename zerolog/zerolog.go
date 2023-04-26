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
	level, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		panic(err)
	}

	// create new instance
	logger := zerolog.New(w).Level(level).With().
		Str("module", module).
		Timestamp().
		Logger()

	zlogger := &ZeroLogger{
		Logger: &logger,
		module: module,
	}

	for _, opt := range opts {
		zlogger.Logger = opt(zlogger.Logger)
	}

	return zlogger
}

func (z *ZeroLogger) With(args ...interface{}) types.Logger {
	newLogger := z.Logger.With().Fields(args).Logger()
	z.Logger = &newLogger
	return z
}

func (z *ZeroLogger) Level(level string) error {
	l, err := zerolog.ParseLevel(level)
	if err != nil {
		return err
	}

	newLogger := z.Logger.Level(l)
	z.Logger = &newLogger

	return nil
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

func (z *ZeroLogger) Trace(msg string, args ...interface{}) {
	z.Logger.Trace().Fields(args).Msg(msg)
}
