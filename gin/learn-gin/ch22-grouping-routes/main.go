package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
http://localhost:8080/v1/users/comments
http://localhost:8080/v2/ping/
*/

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	addUserRoutes(v1)
	addPingRoutes(v1)

	v2 := r.Group("/v2")
	addPingRoutes(v2)

	r.Run()
}

func addPingRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/ping")

	ping.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
}

func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users")
	})
	users.GET("/comments", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users comments")
	})
	users.GET("/pictures", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users pictures")
	})
}
