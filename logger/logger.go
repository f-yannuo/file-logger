package logger

import (
	"github.com/sirupsen/logrus"
)

var ipfsLogger *logrus.Entry

//func AddFields(sysId, sysName, svcId, svcName, procMsg string) {
//	ipfsLogger = logrus.WithFields(logrus.Fields{
//		"sysId":   sysId,
//		"sysName": sysName,
//		"svcId":   svcId,
//		"svcName": svcName,
//		"procMsg": procMsg,
//	})
//}
func AddField(fields logrus.Fields) {
	ipfsLogger = logrus.WithFields(fields)
}
func Error(isCR bool, globalSeqNo, errMsg, svcId string, args ...interface{}) {
	cr := 0
	if isCR {
		cr = 1
	}
	ipfsLogger.WithFields(logrus.Fields{"isCR": cr, "globalSeqNo": globalSeqNo, "errMsg": errMsg, "svcId": svcId}).Error(args...)
}
func Errorf(isCR bool, globalSeqNo, errMsg, svcId, format string, args ...interface{}) {
	cr := 0
	if isCR {
		cr = 1
	}
	ipfsLogger.WithFields(logrus.Fields{"isCR": cr, "globalSeqNo": globalSeqNo, "errMsg": errMsg, "svcId": svcId}).Errorf(format, args...)
}
func Warn(globalSeqNo, warnMsg, svcId string, args ...interface{}) {
	ipfsLogger.WithFields(logrus.Fields{"globalSeqNo": globalSeqNo, "warnMsg": warnMsg, "svcId": svcId}).Warn(args...)
}
func Warnf(globalSeqNo, warnMsg, svcId, format string, args ...interface{}) {
	ipfsLogger.WithFields(logrus.Fields{"globalSeqNo": globalSeqNo, "warnMsg": warnMsg, "svcId": svcId}).Warnf(format, args...)
}
func Info(globalSeqNo, svcId string, args ...interface{}) {
	ipfsLogger.WithFields(logrus.Fields{"globalSeqNo": globalSeqNo, "svcId": svcId}).Info(args...)
}
func Infof(globalSeqNo, svcId, format string, args ...interface{}) {
	ipfsLogger.WithFields(logrus.Fields{"globalSeqNo": globalSeqNo, "svcId": svcId}).Infof(format, args...)
}
func Debug(globalSeqNo, svcId string, args ...interface{}) {
	ipfsLogger.WithFields(logrus.Fields{"globalSeqNo": globalSeqNo, "svcId": svcId}).Debug(args...)
}
func Debugf(globalSeqNo, svcId, format string, args ...interface{}) {
	ipfsLogger.WithFields(logrus.Fields{"globalSeqNo": globalSeqNo, "svcId": svcId}).Debugf(format, args...)
}
