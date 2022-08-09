package log

import (
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/go-angle/angle/config"
	"github.com/go-angle/angle/di"
	"github.com/go-angle/angle/util"
)

func init() {
	di.Provide(newLogger)
}

// jsonLog impl Logger via zerolog
type jsonLog struct {
	logger zerolog.Logger
	ctx    zerolog.Context
}

func newLogger(c *config.Config) Logger {
	ctx := log.With().
		Str("app", c.Name).
		Str("stage", c.Stage).Str("hostname", util.Hostname())

	if c.IsDevelopment() {
		ctx = ctx.Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr}).With()
	}
	return &jsonLog{
		logger: ctx.Logger(),
		ctx:    ctx,
	}
}

// Output set output
func (l *jsonLog) Output(w io.Writer) Logger {
	logger := l.logger.Output(w)
	return &jsonLog{logger: logger, ctx: logger.With()}
}

// WithFields add extra fields
func (l *jsonLog) WithFields(fields Fields) Logger {
	ctx := l.ctx.Fields(fields)
	return &jsonLog{logger: ctx.Logger(), ctx: ctx}
}

// Debugf debug
func (l *jsonLog) Debugf(format string, args ...interface{}) {
	l.logger.Debug().Msgf(format, args...)
}

// Infof info
func (l *jsonLog) Infof(format string, args ...interface{}) {
	l.logger.Info().Msgf(format, args...)
}

// Printf print
func (l *jsonLog) Printf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

// Warnf warnf
func (l *jsonLog) Warnf(format string, args ...interface{}) {
	l.logger.Warn().Msgf(format, args...)
}

// Warningf warn
func (l *jsonLog) Warningf(format string, args ...interface{}) {
	l.Warnf(format, args...)
}

// Errorf error
func (l *jsonLog) Errorf(format string, args ...interface{}) {
	l.logger.Error().Msgf(format, args...)
}

// Fatalf fatalf
func (l *jsonLog) Fatalf(format string, args ...interface{}) {
	l.logger.Fatal().Msgf(format, args...)
}

// Panicf panic
func (l *jsonLog) Panicf(format string, args ...interface{}) {
	l.logger.Panic().Msgf(format, args...)
}

// Debug debug
func (l *jsonLog) Debug(args ...interface{}) {
	l.logger.Debug().Msg(fmt.Sprint(args...))
}

// Info info
func (l *jsonLog) Info(args ...interface{}) {
	l.logger.Info().Msg(fmt.Sprint(args...))
}

// Print print
func (l *jsonLog) Print(args ...interface{}) {
	l.logger.Print(fmt.Sprint(args...))
}

// Warn warn
func (l *jsonLog) Warn(args ...interface{}) {
	l.logger.Warn().Msg(fmt.Sprint(args...))
}

// Warning warn
func (l *jsonLog) Warning(args ...interface{}) {
	l.Warn(args...)
}

// Error error
func (l *jsonLog) Error(args ...interface{}) {
	l.logger.Error().Msg(fmt.Sprint(args...))
}

// Fatal fatal
func (l *jsonLog) Fatal(args ...interface{}) {
	l.logger.Fatal().Msg(fmt.Sprint(args...))
}

// Panic panic
func (l *jsonLog) Panic(args ...interface{}) {
	l.logger.Panic().Msg(fmt.Sprint(args...))
}
