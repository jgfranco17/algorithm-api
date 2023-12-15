package logging

import (
	"github.com/sirupsen/logrus"
)

func GetLogger() *logrus.Logger {
	logger := logrus.New()
	logLevel := logrus.InfoLevel
	logger.SetLevel(logLevel)
	return logger
}
