package logs

import (
	"mysql-agent/config"

	"github.com/sirupsen/logrus"
)

func NewLogger(config *config.Config) *logrus.Logger {
	logger := logrus.New()
	// logger.SetFormatter(&logrus.JSONFormatter{})
	return logger
}
