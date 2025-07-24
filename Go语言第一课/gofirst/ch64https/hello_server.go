package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Println("hello golang \n")
	})
	http.ListenAndServe(":8080", nil)

}
