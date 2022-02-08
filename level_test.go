package traininggolanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()

	logger.Trace("This Is Trace")
	logger.Debug("This Is Debug")
	logger.Info("This Is Info")
	logger.Warn("This Is Warn")
	logger.Error("This Is Error")
}

func TestSetLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel) // Di Set mulai dari trace

	logger.Trace("This Is Trace")
	logger.Debug("This Is Debug")
	logger.Info("This Is Info")
	logger.Warn("This Is Warn")
	logger.Error("This Is Error")
}
