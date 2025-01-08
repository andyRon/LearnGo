package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// 把默认的ASCII格式改为JSON格式
	//log.SetFormatter(&log.JSONFormatter{})
	// 输出stdout而不是默认的stderr，也可以是一个文件
	log.SetOutput(os.Stdout)
	// 只记录警告及以上的日志
	//log.SetLevel(log.WarnLevel)
}

func main() {
	log.WithFields(log.Fields{
		"tool":  "pen",
		"price": "10",
	}).
		Info("The pen price is 10 dollars.")
	//Warn("The pen price is 10 dollars.")
	//Fatal("The pen price is 10 dollars.")

	contextLogger := log.WithFields(log.Fields{
		"common": "这是个字段",
		"other":  "其他想记录的东西",
	})
	contextLogger.Info("这是个带有字段的日志")
}
