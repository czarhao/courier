package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger = logrus.New()

func init() {
	Logger.Formatter = new(logrus.TextFormatter)
	Logger.Formatter.(*logrus.TextFormatter).DisableTimestamp = true
	Logger.Level = logrus.InfoLevel
	Logger.Out = os.Stdout
}
