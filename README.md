# file-logger

# SYNOPSIS

```go

import (
	log "github.com/f-yannuo/file-logger"
	"github.com/f-yannuo/file-logger/formatter"
	"github.com/sirupsen/logrus"
	"time"
)
//使用案例
func init() {

    logLevel:=formatter.Debug//日志等级
    logPath:="./log"//保存日志文件路径
	fields := logrus.Fields{
		"sysId":   "sysId",
		"sysName": "sysName",
		"svcId":   "svcId",
		"svcName": "svcName",
		"procMsg": "procMsg",
	}//拓展名
    maxAge:=10*time.Minute//日志最大保存时间
    rotationTime:=10*time.Minute//日志切割时间
	log.LoggerInit(logLevel, logPath, fields, maxAge, rotationTime)
}
```
# DESCRIPTION

To install, simply issue a `go get`:

```
go get github.com/f-yannuo/file-logger
```
```go

import (
      "github.com/f-yannuo/file-logger/logger"
)
func main()  {
 	logger.Debug("", "debug", "the debug ")
 	logger.Debugf("", "debug", "the debug %s", "f")
 	logger.Info("", "info", "the info ")
 	logger.Infof("", "info", "the info  %s", "f")
 	logger.Error(false, "", "err", "err", "the err ")
 	logger.Errorf(false, "", "err", "err", "the err  %s", "f")
 	logger.Warn("", "warn", "warn", "the warn ")
 	logger.Warnf("", "warn", "warn", "the warn  %s", "f")
}
```