package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 使用 AsciiJSON 生成具有转义的非 ASCII 字符的 ASCII-only JSON
func main() {
	r := gin.Default()
	r.GET("/someJson", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "Go语言",
			"tag":  "<br>",
			"name": "andyron",
		}

		c.AsciiJSON(http.StatusOK, data)
	})
	r.Run()
}
