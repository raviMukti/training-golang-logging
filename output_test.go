package traininggolanglogging

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestOutputFile(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	logger.SetOutput(file)

	logger.Info("Info Ditulis Ke File")
	logger.Warn("Warn Ditulis Ke File")
	logger.Error("Error Ditulis Ke File")
}
