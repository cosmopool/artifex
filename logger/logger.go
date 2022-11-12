package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

var DebugLevel = false
var InfoLevel = true

var prodLogger, pErr = zap.NewProduction()
var debugLogger, dErr = zap.NewDevelopment()

func GetLogger() *zap.SugaredLogger {
	// flushes buffer, if any
	defer prodLogger.Sync()
	defer debugLogger.Sync()

	if DebugLevel {
		if dErr != nil {
			fmt.Println("Could not get debug logger:", dErr)
			os.Exit(3)
		}
		return debugLogger.Sugar()
	} else {
		if pErr != nil {
			fmt.Println("Could not get debug logger:", pErr)
			os.Exit(3)
		}
		return prodLogger.Sugar()
	}
}
