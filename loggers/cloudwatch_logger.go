package loggers

import (
	logrus "github.com/sirupsen/logrus"
)

type cloudwatchLoggerImpl struct {
}
func(log *cloudwatchLoggerImpl) Error(args ...interface{}) {
	logrus.Error(args)

	// here goes the pushes to aws cloud watch
}
func(log *cloudwatchLoggerImpl) Info(args ...interface{}) {
	logrus.Info(args)

	// here goes the pushes to aws cloud watch
}
