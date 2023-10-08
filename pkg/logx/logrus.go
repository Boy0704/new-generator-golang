package logx

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

// InitLogger initializes the logger with the desired configuration.
func InitLogger(logLevel string, logFile string) {
	// Set the log level
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.Fatalf("Invalid log level: %s", logLevel)
	}
	log.SetLevel(level)

	// Log to file
	if logFile != "" {
		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}
		log.SetOutput(file)
	} else {
		// Log to console if no file is provided
		log.SetOutput(os.Stdout)
	}

	// Customize log format
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

// GetLogger returns the instance of the logger.
func GetLogger() *logrus.Logger {
	return log
}
