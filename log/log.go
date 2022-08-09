package log

import (
	"io"

	"github.com/go-angle/angle/di"
)

func init() {
	// initialize global instance
	di.Invoke(func(l Logger) {
		g = l
	})
}

var g Logger

// Fields alias to logrus.Fields
type Fields map[string]interface{}

// Logger interface of log
type Logger interface {
	Printf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Print(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	WithFields(Fields) Logger

	// Output set output writer
	Output(io.Writer) Logger
}

// Debugf debug
func Debugf(format string, args ...interface{}) {
	g.Debugf(format, args...)
}

// Infof info
func Infof(format string, args ...interface{}) {
	g.Infof(format, args...)
}

// Printf print
func Printf(format string, args ...interface{}) {
	g.Printf(format, args...)
}

// Warnf warn
func Warnf(format string, args ...interface{}) {
	g.Warnf(format, args...)
}

// Warningf warning
func Warningf(format string, args ...interface{}) {
	g.Warningf(format, args...)
}

// Errorf error
func Errorf(format string, args ...interface{}) {
	g.Errorf(format, args...)
}

// Fatalf fatal
func Fatalf(format string, args ...interface{}) {
	g.Fatalf(format, args...)
}

// Panicf panic
func Panicf(format string, args ...interface{}) {
	g.Panicf(format, args...)
}

// Debug debug
func Debug(args ...interface{}) {
	g.Debug(args...)
}

// Info info
func Info(args ...interface{}) {
	g.Info(args...)
}

// Print print
func Print(args ...interface{}) {
	g.Print(args...)
}

// Warn warn
func Warn(args ...interface{}) {
	g.Warn(args...)
}

// Warning warn
func Warning(args ...interface{}) {
	g.Warning(args...)
}

// Error error
func Error(args ...interface{}) {
	g.Error(args...)
}

// Fatal fatal
func Fatal(args ...interface{}) {
	g.Fatal(args...)
}

// Panic panic
func Panic(args ...interface{}) {
	g.Panic(args...)
}

// Output set log output
func Output(w io.Writer) Logger {
	return g.Output(w)
}

// WithFields add extra fields
func WithFields(fields Fields) Logger {
	return g.WithFields(Fields(fields))
}
