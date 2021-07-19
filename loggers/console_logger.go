package loggers

import (
	logrus "github.com/sirupsen/logrus"
)

type consoleLoggerImpl struct {
}
func(log *consoleLoggerImpl) Error(args ...interface{}) {
	logrus.Error(args)
}
func(log *consoleLoggerImpl) Info(args ...interface{}) {
	logrus.Info(args)
}