package main

import (
	"logdemo/mylogger"
)

func main() {
	log := mylogger.NewLog("info")
	log.Debug("这是debug级别的日志")
	log.Info("这是info级别的日志")
	log.Warning("这是warning级别的日志")
	log.Error("这是error级别的日志")
}
