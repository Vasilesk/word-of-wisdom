package zero

import (
	"os"
	"time"

	"github.com/rs/zerolog"

	"github.com/vasilesk/word-of-wisdom/pkg/logger"
)

type zero struct {
	app string

	z *zerolog.Logger

	err  error
	data map[string]interface{}
}

//nolint:reassign
func New(app string) logger.Logger {
	zerolog.LevelFieldName = "lvl"
	zerolog.MessageFieldName = "msg"
	zerolog.ErrorFieldName = "err"
	zerolog.TimestampFieldName = "t"
	zerolog.TimeFieldFormat = time.RFC3339

	z := zerolog.New(os.Stdout).With().Timestamp().Logger()

	l := newInstance(app, &z, nil, nil)

	return l
}

func newInstance(
	app string,
	z *zerolog.Logger,
	err error,
	data map[string]interface{},
) *zero {
	return &zero{
		app:  app,
		z:    z,
		err:  err,
		data: data,
	}
}

func (z *zero) Infof(format string, v ...interface{}) {
	z.msgf(zerolog.InfoLevel, format, v...)
}

func (z *zero) Warnf(format string, v ...interface{}) {
	z.msgf(zerolog.WarnLevel, format, v...)
}

func (z *zero) Errorf(format string, v ...interface{}) {
	z.msgf(zerolog.ErrorLevel, format, v...)
}

func (z *zero) Fatalf(format string, v ...interface{}) {
	z.msgf(zerolog.FatalLevel, format, v...)
	os.Exit(1)
}

func (z *zero) WithError(err error) logger.Logger {
	return newInstance(z.app, z.z, err, z.data)
}

func (z *zero) WithData(data map[string]interface{}) logger.Logger {
	return newInstance(z.app, z.z, z.err, data)
}

func (z *zero) msgf(level zerolog.Level, format string, v ...interface{}) {
	ev := z.z.WithLevel(level)

	ev.Str("app", z.app)

	if z.data != nil {
		ev.Interface("data", z.data)
	}

	if z.err != nil {
		ev = ev.Err(z.err)
	}

	ev.Msgf(format, v...)
}
