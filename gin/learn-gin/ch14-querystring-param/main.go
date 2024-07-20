package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// curl "localhost:8080/welcome?firstname=Andy&lastname=Ron"

func main() {
	r := gin.Default()

	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // c.Request.URL.Query().Get("lastname") 的一种快捷方式

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	r.Run()
}
