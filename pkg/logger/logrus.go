package logger

import "github.com/sirupsen/logrus"

type logrusLogger struct {
	logger logrus.Logger
}

// NewLogger creates a new instance of logrus logger
func NewLogger(prefix string) Logger {
	logger := logrus.New()
	logger.WithField("prefix", prefix)
	return &logrusLogger{
		logger: *logger,
	}
}

// Info logs on "info" level
func (l logrusLogger) Info(v string) {
	l.logger.Info(v)
}

// Info logs on "info" level with fields
func (l logrusLogger) InfoWithFields(v string, fields LoggerFields) {
	l.logger.WithFields(logrus.Fields(fields)).Info(v)
}

// Info logs on "info" level with error
func (l logrusLogger) InfoError(v string, err error) {
	l.logger.WithError(err).Info(v)
}
