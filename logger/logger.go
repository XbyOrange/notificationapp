package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"time"
	"github.com/getsentry/raven-go"
)


func Init(o io.Writer, level logrus.Level){
	logrus.SetOutput(o)
	logrus.SetLevel(level)
}

func Print(level, service, component, traceid, code, message string, errTrace ...error ) {
	fields := logrus.Fields{
		"@timestamp":  time.Now().UTC().Format(time.RFC3339),
		"templates":     service,
		"component": component,
		"traceid": traceid,
		"code": code,
	}
	switch level {
	case "INFO":
		logrus.WithFields(fields).Info(message)
	case "WARN":
		logrus.WithFields(fields).Warn(message)
	case "ERROR":
		logrus.WithFields(fields).Error(message)
	case "PANIC":
		logrus.WithFields(fields).Panic(message)
	case "FATAL":
		logrus.WithFields(fields).Fatal(message)
	default:
		logrus.WithFields(fields).Debug(message)
	}
}

func GetLevel(level string)logrus.Level{
	switch level{
	case "INFO":
		return logrus.InfoLevel
	case "WARN":
		return logrus.WarnLevel
	case "ERROR":
		return logrus.ErrorLevel
	case "PANIC":
		return logrus.PanicLevel
	case "FATAL":
		return logrus.FatalLevel
	}
	return logrus.InfoLevel
}
