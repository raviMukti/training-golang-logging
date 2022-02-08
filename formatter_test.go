package traininggolanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestFormatterJSON(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Info JSON Formatter")
	logger.Warn("Warn JSON Formatter")
	logger.Error("Error JSON Formatter")
}
