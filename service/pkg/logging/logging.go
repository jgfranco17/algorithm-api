package logging

import (
	"context"

	"github.com/sirupsen/logrus"
)

func GetLogger(ctx context.Context) *logrus.Logger {
	logger := logrus.New()
	logLevel := logrus.InfoLevel
	logger.SetLevel(logLevel)
	logger.WithFields(logrus.Fields{
		"Section": ctx.Value("section"),
	})
	return logger
}
