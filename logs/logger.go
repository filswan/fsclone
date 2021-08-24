package logs

import (
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

var Log *log.Logger

func InitLogger() {
	if Log != nil {
		return
	}

	Log = log.New()
	Log.SetLevel(log.DebugLevel)
	formatter := &log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		FullTimestamp:   true,
	}
	Log.SetReportCaller(true)
	Log.SetFormatter(formatter)
	pathMap := lfshook.PathMap{
		log.InfoLevel:  "./logs/info.log",
		log.WarnLevel:  "./logs/warn.log",
		log.ErrorLevel: "./logs/error.log",
		log.FatalLevel: "./logs/error.log",
		log.PanicLevel: "./logs/error.log",
	}
	Log.Hooks.Add(lfshook.NewHook(
		pathMap,
		formatter,
	))
	Log.WriterLevel(log.InfoLevel)
}
func GetLogger() *log.Logger {
	return Log
}
