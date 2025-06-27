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
		// curl 输出：{"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		/*
			浏览器输出：
			{
				"lang": "Go语言",
				"name": "andyron",
				"tag": "<br>"
			}
			因为AsciiJSON 会将响应中的非 ASCII 字符（如中文、Emoji）转换为 Unicode 转义序列（例如 \u8bed\u8a00），确保输出内容仅包含 ASCII 字符集。
			但浏览器会自动将 Unicode 转义序列还原为原始字符。
		*/
	})
	r.Run()
}
