package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// TODO
func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)

			dst := "./" + file.Filename
			c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	r.Run()
}

/*
curl -X POST http://localhost:8080/upload \
  -F "file=@/Users/andyron/Downloads/gopher.png" \
  -F "file=@/Users/andyron/Downloads/gopher2.png" \
  -H "Content-Type: multipart/form-data"
*/
