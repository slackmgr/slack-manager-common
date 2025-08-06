package common

import (
	"io"
)

type NoopLogger struct{}

func (l *NoopLogger) Debug(msg string) {
}

func (l *NoopLogger) Debugf(format string, args ...any) {
}

func (l *NoopLogger) Info(msg string) {
}

func (l *NoopLogger) Infof(format string, args ...any) {
}

func (l *NoopLogger) Error(msg string) {
}

func (l *NoopLogger) Errorf(format string, args ...any) {
}

func (l *NoopLogger) HttpLoggingHandler() io.Writer {
	return nil
}

func (l *NoopLogger) WithField(key string, value any) Logger { //nolint:ireturn
	return l
}

func (l *NoopLogger) WithFields(fields map[string]any) Logger { //nolint:ireturn
	return l
}
