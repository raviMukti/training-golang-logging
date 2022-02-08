package traininggolanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()

	logger.Info("Hello World")
}

func TestWithField(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "ravi").WithField("name", "Ravi").Info("Hello World")
}

func TestFields(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "ravi",
		"name":     "Ravi",
	}).Info("Hello World")
}
