package loggers

import (
	configService "simpl.com/configs"
)

// exposed logger object
var Logger = getLogger()

/*
 definitions
*/
type loggers interface {
	Error(args ...interface{})
	Info(args ...interface{})
}

var consoleLogger = consoleLoggerImpl {
}
var cloudwatchLogger = cloudwatchLoggerImpl {
}

// factory function for logger objects
func getLogger() loggers {

	environment := configService.Configs.Environment
	switch environment {
		case "delelopment" :
			return &consoleLogger
		case "test" :
			return &cloudwatchLogger
		default:
			return &consoleLogger
	}
}

