package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

// HTML 渲染

func main() {
	//t1()
	//t2()
	//t3()
	t4()
}

func t1() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8080")
}

func t2() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "posts",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "users",
		})
	})
	r.Run()
}

// 自定义模板渲染器
func t3() {
	r := gin.Default()
	html := template.Must(template.ParseFiles("file1", "file2")) // TODO
	r.SetHTMLTemplate(html)
	r.Run()
}

func t4() {
	r := gin.Default()
	r.Delims("{[{", "}]}") // 自定义分隔符
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	r.LoadHTMLFiles("templates/raw.tmpl")

	r.GET("/raw", func(c *gin.Context) {
		c.HTML(http.StatusOK, "raw.tmpl", map[string]interface{}{
			"now": time.Date(2024, 07, 01, 0, 0, 0, 0, time.UTC),
		})
	})
	r.Run()
}
func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}
