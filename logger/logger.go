package logger

import "go.uber.org/zap"

var DebugLevel = false
var InfoLevel = true

var prodLogger, pErr = zap.NewProduction()
var debugLogger, dErr = zap.NewDevelopment()

func GetLogger() (*zap.SugaredLogger, error) {
	// flushes buffer, if any
	defer prodLogger.Sync()
	defer debugLogger.Sync()

	if DebugLevel {
		return debugLogger.Sugar(), dErr
	} else {
		return prodLogger.Sugar(), pErr
	}
}
