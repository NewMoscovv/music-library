package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var Logger *logrus.Logger

func Init(logLevel string) {
	Logger = logrus.New()

	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	Logger.SetOutput(os.Stdout)

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		Logger.SetLevel(logrus.DebugLevel)
	} else {
		Logger.SetLevel(level)
	}
	Logger.Info("Логгер инициализирован")
}

func Debug(message string, fields map[string]interface{}) {
	Logger.WithFields(fields).Debug(message)
}

func Info(message string, fields map[string]interface{}) {
	Logger.WithFields(fields).Info(message)
}

func Warn(message string, fields map[string]interface{}) {
	Logger.WithFields(fields).Warn(message)
}

func Error(message string, fields map[string]interface{}) {
	Logger.WithFields(fields).Error(message)
}

func Fatal(message string, fields map[string]interface{}) {
	Logger.WithFields(fields).Fatal(message)
}
