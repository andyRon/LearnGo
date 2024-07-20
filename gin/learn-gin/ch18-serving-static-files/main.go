package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// TODO
func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.StaticFS("/more_static", http.Dir("my_file_system"))
	r.StaticFile("/favicon.ico", "./resources/favicon.ico")

	r.Run()
}
