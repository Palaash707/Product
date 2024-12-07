package logging

import (
	"go.uber.org/zap"
)

// Logger is the global zap logger instance
var Logger *zap.Logger

// InitLogger initializes the zap logger
func InitLogger() {
    var err error
    Logger, err = zap.NewProduction()  // or zap.NewDevelopment() for less structured output

    if err != nil {
        panic(err)
    }
}
