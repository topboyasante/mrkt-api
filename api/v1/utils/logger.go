package utils

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type CustomFormatter struct {
	logrus.Formatter
}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var file string
	if entry.HasCaller() {
		file = fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
	}
	entry.Message = fmt.Sprintf("%s %s", entry.Message, file)

	return f.Formatter.Format(entry)
}

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.SetReportCaller(true)
	logger.SetFormatter(&CustomFormatter{
		Formatter: &logrus.JSONFormatter{},
	})
	logger.SetLevel(logrus.InfoLevel)
}

func Logger() *logrus.Logger {
	return logger
}
