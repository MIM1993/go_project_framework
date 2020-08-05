package internal

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strings"
	"time"
)

// initialize log
func LogSet(logLevel logrus.Level, pretty bool, logFile string) error {
	logger := logrus.StandardLogger()

	logger.SetOutput(os.Stdout)
	logger.SetReportCaller(true)
	logger.SetLevel(logrus.Level(logLevel))
	formatter := &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.SplitAfterN(f.File, "go_project_framework/", 2)
			filename := s[len(s)-1]
			return "", fmt.Sprintf("%s:%d", filename, f.Line)
		},
		PrettyPrint: pretty,
	}
	logger.SetFormatter(formatter)

	if len(logFile) > 0 {
		logfile, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}

		logger.SetOutput(logfile)
	}

	return nil
}
