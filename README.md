# file-logger

# SYNOPSIS

```go

import (
	log "github.com/f-yannuo/file-logger"
	"github.com/f-yannuo/file-logger/formatter"
	"github.com/f-yannuo/file-logger/logger"
	"github.com/sirupsen/logrus"
	"time"
)
//使用案例
func init() {
	fields := logrus.Fields{
		"sysId":   "sysId",
		"sysName": "sysName",
		"svcId":   "svcId",
		"svcName": "svcName",
		"procMsg": "procMsg",
	}//拓展名
	log.LoggerInit(formatter.Debug,//日志等级
    		"./log", //log path
    		fields,//拓展名
    		10*time.Minute,//日志最大保存时间
    		2*time.Minute,//日志切割时间
    	)
}