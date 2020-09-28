package log

import (
	"github.com/airingone/config"
	"testing"
)

//log test
func TestLog(t *testing.T) {
	config.InitConfig()
	conf := config.GetLogConfig("log")
	InitLog(conf)

	Error("error:%d", 1000)
	logHander := NewLogHandler()
	logHander.SetRequestId("123456")
	logHander.Debug("debug:%+v", conf)

	//t.Logf("server.name: %s", serverName)
}
