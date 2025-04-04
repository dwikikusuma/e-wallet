package bootstrap

import "github.com/sirupsen/logrus"

func SetupLogger() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	log.Info("Logger initiated using log rus")
	return log
}
