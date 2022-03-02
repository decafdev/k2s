package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/techdecaf/k2s/v2/pkg/config"
)

type LoggerI interface {
	Print(...interface{})
	Fatal(...interface{})
}

// NewLogger function description
func NewLogger(config *config.ConfigService) *logrus.Entry {
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
