package log

import (
	"fmt"
	"github.com/airingone/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"runtime"
	"time"
)

var RUSLOG *logrus.Logger

func InitLog(conf config.ConfigLog) {
	RUSLOG = logrus.New()
	timeStr := time.Now().Format("2006-01-02")
	logger := &lumberjack.Logger{
		Filename:   conf.Path + "log-" + timeStr + ".log",
		MaxSize:    int(conf.MaxSize),    //MB
		MaxBackups: int(conf.MaxBackups), //file count
		MaxAge:     int(conf.MaxAge),     //days
		Compress:   conf.Compress,        //disabled by default
	}

	if conf.Level == "error" {
		RUSLOG.SetLevel(logrus.ErrorLevel)
	} else if conf.Level == "debug" {
		RUSLOG.SetLevel(logrus.DebugLevel)
	} else if conf.Level == "info" {
		RUSLOG.SetLevel(logrus.InfoLevel)
	} else if conf.Level == "warn" {
		RUSLOG.SetLevel(logrus.WarnLevel)
	} else if conf.Level == "trace" {
		RUSLOG.SetLevel(logrus.TraceLevel)
	} else if conf.Level == "panic" {
		RUSLOG.SetLevel(logrus.PanicLevel)
	} else if conf.Level == "fatal" {
		RUSLOG.SetLevel(logrus.FatalLevel)
	} else {
		RUSLOG.SetLevel(logrus.InfoLevel)
	}

	RUSLOG.SetFormatter(&logrus.JSONFormatter{})
	//RUSLOG.SetFormatter(&logrus.TextFormatter{})
	RUSLOG.SetOutput(logger)
}

func Error(format string, msg ...interface{}) {
	if RUSLOG != nil {
		_, file, line, _ := runtime.Caller(1)
		strMsg := fmt.Sprintf(format, msg...)
		RUSLOG.WithFields(logrus.Fields{"file": file, "line": line}).Error(strMsg)
	}
}

func Debug(format string, msg ...interface{}) {
	if RUSLOG != nil {
		_, file, line, _ := runtime.Caller(1)
		strMsg := fmt.Sprintf(format, msg...)
		RUSLOG.WithFields(logrus.Fields{"file": file, "line": line}).Debug(strMsg)
	}
}

func Info(format string, msg ...interface{}) {
	if RUSLOG != nil {
		_, file, line, _ := runtime.Caller(1)
		strMsg := fmt.Sprintf(format, msg...)
		RUSLOG.WithFields(logrus.Fields{"file": file, "line": line}).Info(strMsg)
	}
}

func Warn(format string, msg ...interface{}) {
	if RUSLOG != nil {
		_, file, line, _ := runtime.Caller(1)
		strMsg := fmt.Sprintf(format, msg...)
		RUSLOG.WithFields(logrus.Fields{"file": file, "line": line}).Warn(strMsg)
	}
}

func Trace(format string, msg ...interface{}) {
	if RUSLOG != nil {
		_, file, line, _ := runtime.Caller(1)
		strMsg := fmt.Sprintf(format, msg...)
		RUSLOG.WithFields(logrus.Fields{"file": file, "line": line}).Trace(strMsg)
	}
}

func Panic(format string, msg ...interface{}) {
	if RUSLOG != nil {
		_, file, line, _ := runtime.Caller(1)
		strMsg := fmt.Sprintf(format, msg...)
		RUSLOG.WithFields(logrus.Fields{"file": file, "line": line}).Panic(strMsg)
	}
}

func Fatal(format string, msg ...interface{}) {
	if RUSLOG != nil {
		_, file, line, _ := runtime.Caller(1)
		strMsg := fmt.Sprintf(format, msg...)
		RUSLOG.WithFields(logrus.Fields{"file": file, "line": line}).Fatal(strMsg)
	}
}

func PanicTrack() {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	stackInfo := fmt.Sprintf("%s", buf[:n])
	Error("panic stack info %s", stackInfo)
}

type LogHandler struct {
	Logger    *logrus.Logger
	RequestId string //请求唯一Id
}

func NewLogHandler() *LogHandler {
	var handle LogHandler
	handle.Logger = RUSLOG

	return &handle
}

func (h *LogHandler) SetRequestId(requestId string) {
	h.RequestId = requestId
}

func (h *LogHandler) Error(format string, msg ...interface{}) {
	if h.Logger != nil {
		_, file, line, _ := runtime.Caller(1)
		strMsg := fmt.Sprintf(format, msg...)
		h.Logger.WithFields(logrus.Fields{"requestId": h.RequestId, "file": file, "line": line}).Error(strMsg)
	}
}

func (h *LogHandler) Debug(format string, msg ...interface{}) {
	if h.Logger != nil {
		_, file, line, _ := runtime.Caller(1)
		strMsg := fmt.Sprintf(format, msg...)
		h.Logger.WithFields(logrus.Fields{"requestId": h.RequestId, "file": file, "line": line}).Debug(strMsg)
	}
}

func (h *LogHandler) Info(format string, msg ...interface{}) {
	if h.Logger != nil {
		_, file, line, _ := runtime.Caller(1)
		strMsg := fmt.Sprintf(format, msg...)
		h.Logger.WithFields(logrus.Fields{"requestId": h.RequestId, "file": file, "line": line}).Info(strMsg)
	}
}

func (h *LogHandler) Warn(format string, msg ...interface{}) {
	if h.Logger != nil {
		_, file, line, _ := runtime.Caller(1)
		strMsg := fmt.Sprintf(format, msg...)
		h.Logger.WithFields(logrus.Fields{"requestId": h.RequestId, "file": file, "line": line}).Warn(strMsg)
	}
}

func (h *LogHandler) Trace(format string, msg ...interface{}) {
	if h.Logger != nil {
		_, file, line, _ := runtime.Caller(1)
		strMsg := fmt.Sprintf(format, msg...)
		h.Logger.WithFields(logrus.Fields{"requestId": h.RequestId, "file": file, "line": line}).Trace(strMsg)
	}
}

func (h *LogHandler) Panic(format string, msg ...interface{}) {
	if h.Logger != nil {
		_, file, line, _ := runtime.Caller(1)
		strMsg := fmt.Sprintf(format, msg...)
		h.Logger.WithFields(logrus.Fields{"requestId": h.RequestId, "file": file, "line": line}).Panic(strMsg)
	}
}

func (h *LogHandler) Fatal(format string, msg ...interface{}) {
	if h.Logger != nil {
		_, file, line, _ := runtime.Caller(1)
		strMsg := fmt.Sprintf(format, msg...)
		h.Logger.WithFields(logrus.Fields{"requestId": h.RequestId, "file": file, "line": line}).Fatal(strMsg)
	}
}
