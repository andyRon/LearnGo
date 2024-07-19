package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// TODO
func main() {
	r := gin.Default()
	r.GET("/jsonp", func(c *gin.Context) {
		data := map[string]interface{}{
			"name": "tom",
		}

		c.JSONP(http.StatusOK, data)
	})
	r.Run()
}

/**
/jsonp?callback=x
x({"name":"tom"});
*/
