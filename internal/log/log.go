package log

import "github.com/sirupsen/logrus"

func init() {
	// 判断用户是否输入准确,默认是info
	logrus.SetLevel(logrus.TraceLevel)
}

func Debug(v ...interface{}) {
	logrus.Debug(v)
}

func Info(v ...interface{}) {
	logrus.Info(v)
}

func Warn(v ...interface{}) {
	logrus.Warn(v)
}

func Error(v ...interface{}) {
	logrus.Error(v)
}
