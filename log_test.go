package log

import (
	"github.com/airingone/config"
	"testing"
)

func TestLog(t *testing.T) {
	config.InitConfig()
	conf := config.GetLogConfig("log")
	InitLog(conf)

	Error("errorf:%d", 1000)
	logHander := NewLogHander("123456789")
	logHander.Debug("debugf:%+v", conf)

	//t.Logf("server.name: %s", serverName)
}
