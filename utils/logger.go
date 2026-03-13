package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

func SetupLogger() *logrus.Logger {
	log := logrus.New()
	log.SetOutput(os.Stdout)

	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
		PadLevelText:  true,
	})

	return log
}
