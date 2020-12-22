package file_logger

import (
	"github.com/f-yannuo/file-logger/formatter"
	"github.com/f-yannuo/file-logger/logger"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func LoggerInit(logLevel formatter.Level, logPath string, fields logrus.Fields, maxAge, rotationTime time.Duration) {
	//add log fields
	logger.AddField(fields)
	// set log level
	formatter.ConfigLoglevel(logLevel)
	// set log format
	formatter.ConfigLogFormatter()
	// set log output
	logrus.SetOutput(os.Stdout)
	formatter.ConfigLoggerWithPath(logPath, maxAge, rotationTime)
}
