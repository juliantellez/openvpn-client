package logger_test

import (
	"testing"

	"github.com/juliantellez/openvpn-client/shared/logger"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	err := logger.New("warn", "text")
	assert.NoError(t, err)

	// Text format
	assert.Equal(t, logrus.WarnLevel, logrus.StandardLogger().Level)
	_, ok := logrus.StandardLogger().Formatter.(*logrus.TextFormatter)
	assert.True(t, ok)

	// Default format
	err = logger.New("warn", "json")
	assert.NoError(t, err)
	_, ok = logrus.StandardLogger().Formatter.(*logrus.JSONFormatter)
	assert.True(t, ok)
}

func TestNewLogger_Errors(t *testing.T) {
	err := logger.New("", "json")
	assert.EqualError(t, err, logger.ErrorLoggerLogLevel)

	err = logger.New("foo", "json")
	assert.EqualError(t, err, logger.ErrorLoggerLogLevelParse)
}
