package examples

import (
	stdLog "log"
	"os"
)

type ExampleLogger struct {
	debug stdLog.Logger
	info  stdLog.Logger
	warn  stdLog.Logger
	error stdLog.Logger
}

func NewExampleLogger() *ExampleLogger {
	return &ExampleLogger{
		debug: *stdLog.New(os.Stdout, "[DEBUG] ", stdLog.Ldate|stdLog.Ltime|stdLog.Lmsgprefix),
		info:  *stdLog.New(os.Stdout, "[INFO] ", stdLog.Ldate|stdLog.Ltime|stdLog.Lmsgprefix),
		warn:  *stdLog.New(os.Stdout, "[WARN] ", stdLog.Ldate|stdLog.Ltime|stdLog.Lmsgprefix),
		error: *stdLog.New(os.Stdout, "[ERROR] ", stdLog.Ldate|stdLog.Ltime|stdLog.Lmsgprefix),
	}
}

func (l *ExampleLogger) Debug(msg string, args ...any) {
	l.debug.Printf(msg, args...)
}

func (l *ExampleLogger) Info(msg string, args ...any) {
	l.info.Printf(msg, args...)
}

func (l *ExampleLogger) Warn(msg string, args ...any) {
	l.warn.Printf(msg, args...)
}

func (l *ExampleLogger) Error(msg string, args ...any) {
	l.error.Printf(msg, args...)
}
