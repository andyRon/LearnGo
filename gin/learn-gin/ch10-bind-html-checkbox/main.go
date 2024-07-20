package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// TODO

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./*")
	r.GET("/form", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html", gin.H{})
	})
	r.Run()

}

type myForm struct {
	Colors []string `form:"colors[]"`
}

func formHandler(c *gin.Context) {
	var fakeForm myForm
	c.ShouldBind(&fakeForm)
	c.JSON(http.StatusOK, gin.H{"color": fakeForm.Colors})
}
