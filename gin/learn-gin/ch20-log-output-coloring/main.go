package main

import "github.com/gin-gonic/gin"

func main() {
	// 强制日志颜色化 默认
	gin.ForceConsoleColor()
	// 禁止日志的颜色
	//gin.DisableConsoleColor()

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.Run()
}
