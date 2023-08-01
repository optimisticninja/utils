package logger

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

func GetLogger(f logrus.Fields) *logrus.Entry {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors:          false,
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		TimestampFormat:        "2006-01-02T15:04:05.000Z0700",
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			pathParts := strings.Split(f.File, "/")
			filename := pathParts[len(pathParts)-1]
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})
	return logger.WithFields(f)
}
