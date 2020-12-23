package formatter

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

const (
	Panic  = "panic"
	Fatal  = "fatal"
	Error  = "error"
	Warn   = "warn"
	Info   = "info"
	Debug  = "debug"
)

func ConfigLoglevel(logLevel string) {
	switch logLevel {
	case Panic:
		log.SetLevel(log.PanicLevel)
	case Fatal:
		log.SetLevel(log.FatalLevel)
	case Error:
		log.SetLevel(log.ErrorLevel)
	case Warn:
		log.SetLevel(log.WarnLevel)
	case Info:
		log.SetLevel(log.InfoLevel)
	case Debug:
		log.SetLevel(log.DebugLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}
func ConfigLogFormatter() {
	customTextFormatter := &SpdbFormatter{}
	customTextFormatter.TimestampFormat = "2006-01-02 15:04:05.999"
	customTextFormatter.FullTimestamp = true
	customTextFormatter.FieldMap = FieldMap{
		log.FieldKeyTime:  "timeStamp",
		log.FieldKeyLevel: "level",
		log.FieldKeyMsg:   "message"}
	customTextFormatter.Splitter = ","
	log.SetFormatter(customTextFormatter)
}

func ConfigLoggerWithPath(logPath string, maxAge, rotationTime time.Duration) {

	if err := os.MkdirAll(logPath, os.ModePerm); err != nil {
		log.Errorf("config local file system logger error, %+v", errors.WithStack(err))
		return
	}
	debugLfHook := newHook(logPath, Debug, maxAge, rotationTime)
	infoLfHook := newHook(logPath, Info, maxAge, rotationTime)
	errLfHook := newHook(logPath, Error, maxAge, rotationTime)
	log.AddHook(debugLfHook)
	log.AddHook(infoLfHook)
	log.AddHook(errLfHook)
}

//"/%Y-%m-%d/"
func newHook(logPath ,level string , maxAge, rotationTime time.Duration) *lfshook.LfsHook {
	var (
		writerMap lfshook.WriterMap
		ti        = //""
		"/%Y-%m-%d-%H-%M/"
	)
	fileName := ti + string(level) + "Writer.log"
	writerPath := path.Join(logPath, ti+string(level)+"Writer.log")
	writer, _ := rotatelogs.New(
		logPath+fileName,
		rotatelogs.WithLinkName(writerPath),       // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	switch level {
	case Error:
		writerMap = lfshook.WriterMap{
			log.ErrorLevel: writer,
			log.FatalLevel: writer,
			log.PanicLevel: writer,
		}
	case Info:
		writerMap = lfshook.WriterMap{
			log.InfoLevel: writer,
			log.WarnLevel: writer,
		}
	case Debug:
		writerMap = lfshook.WriterMap{
			log.TraceLevel: writer,
			log.DebugLevel: writer,
			log.InfoLevel:  writer,
			log.WarnLevel:  writer,
			log.ErrorLevel: writer,
			log.FatalLevel: writer,
			log.PanicLevel: writer,
		}
	default:
		if level == "" {
			return newHook(logPath, Debug, maxAge, rotationTime)
		}
	}
	return lfshook.NewHook(writerMap, &SpdbFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999",
		FullTimestamp:   true,
		FieldMap: FieldMap{
			log.FieldKeyTime:  "timeStamp",
			log.FieldKeyLevel: "level",
			log.FieldKeyMsg:   "message"},
		Splitter: ",",
	})
}
