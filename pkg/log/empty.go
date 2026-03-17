package log

type EmptyLogger struct {
}

func (l *EmptyLogger) Debug(msg string, args ...any) {
	//noop
}

func (l *EmptyLogger) Info(msg string, args ...any) {
	//noop
}

func (l *EmptyLogger) Warn(msg string, args ...any) {
	//noop
}

func (l *EmptyLogger) Error(msg string, args ...any) {
	//noop
}
