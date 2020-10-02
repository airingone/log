# 日记文件组件
## 1.组件描述
log用于写日志。

## 2.如何使用
```
import (
    "github.com/airingone/config"
    "github.com/airingone/log"
)

func main() {
    config.InitConfig()                     //进程启动时调用一次初始化配置文件，配置文件名为config.yml，目录路径为../conf/或./
    log.InitLog(config.GetLogConfig("log")) //进程启动时调用一次初始化日志

    #打印日志
    log.Error("error: %d", 10000)

    #设置requestid方式打印日志
    logHander := log.NewLogHandler()
    logHander.SetRequestId
    logHander.Debug("debug: %s", "test")
}
```
更多使用请参考[测试例子](https://github.com/airingone/log/blob/master/log_test.go)
