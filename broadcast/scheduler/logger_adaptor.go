package scheduler

import "github.com/sirupsen/logrus"

type loggerAdaptor struct {
	logger *logrus.Logger
}

func newLoggerAdaptor(logger *logrus.Logger) *loggerAdaptor {
	return &loggerAdaptor{logger}
}

func (l *loggerAdaptor) Error(err error, msg string, keysAndValues ...interface{}) {
	l.logger.Error(err)
}

func (l *loggerAdaptor) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Info(msg)
}
