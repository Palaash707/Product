package logging

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is the global zap logger instance
var Logger *zap.Logger

// InitLogger initializes the zap logger with log rotation using lumberjack
func InitLogger() {
    // Create a lumberjack logger for log rotation
    writer := &lumberjack.Logger{
        Filename:   "app.log",   // Path to the log file
        MaxSize:    10,          // Maximum size in MB before the log file is rotated
        MaxBackups: 3,           // Number of backups to keep
        MaxAge:     28,          // Maximum number of days to retain old log files
    }

    // Create a zapcore.WriteSyncer to wrap the lumberjack writer
    fileSyncer := zapcore.AddSync(writer)

    // Set the encoder for the logs (you can use JSON or human-readable format)
    encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())  // For human-readable format


    // Create a custom core to handle the log writing to the file
    core := zapcore.NewCore(
        encoder,
        fileSyncer,
        zap.InfoLevel, // Set log level to Info (you can adjust based on your need)
    )

    // Create the logger using the custom core
    Logger = zap.New(core)

    // Optional: Log initialization message to verify setup
    Logger.Info("Logger initialized successfully with rotation.")
}
