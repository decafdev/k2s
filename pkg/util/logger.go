package util

import (
	"github.com/sirupsen/logrus"
)

// NewLogger function description
func NewLogger(config *Config) *logrus.Entry {
	// set logging defaults
	if config.LOGGER_PRETTY_PRINT == "true" {
		logrus.SetFormatter(&logrus.TextFormatter{})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	switch config.LOGGER_LEVEL {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	default:
		logrus.SetLevel(logrus.ErrorLevel)
	}

	return logrus.WithFields(logrus.Fields{
		"service": config.SERVICE_NAME,
		"version": config.VERSION,
		"env":     config.ENVIRONMENT,
	})
}
