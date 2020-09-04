package logger

import (
	"errors"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	// ErrorLoggerLogLevel 1
	ErrorLoggerLogLevel = "[ Logger ] log level is empty"
	// ErrorLoggerLogLevelParse 2
	ErrorLoggerLogLevelParse = "[ Logger ] failed to parse log level"
)

// Logger structure
type Logger struct{}

// New creates a new logger instance
func New(level, format string) error {
	if level == "" {
		return errors.New(ErrorLoggerLogLevel)
	}

	logrusLevel, err := logrus.ParseLevel(level)
	if err != nil {
		return errors.New(ErrorLoggerLogLevelParse)
	}
	logrus.SetLevel(logrusLevel)

	logrusFormat := strings.ToLower(format)

	switch logrusFormat {
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{})
	default:
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
	return nil
}
