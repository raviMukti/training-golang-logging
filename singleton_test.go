package traininggolanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T) {
	logrus.Info("Hello Info")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Warn("Hello Warn")
}
