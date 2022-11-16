package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

func New(showDebugLevel bool) *zap.SugaredLogger {
	var prodLogger, pErr = zap.NewProduction()
	var debugLogger, dErr = zap.NewDevelopment()

	// flushes buffer, if any
	defer prodLogger.Sync()
	defer debugLogger.Sync()

	if showDebugLevel {
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
