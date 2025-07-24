package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello https"))
		fmt.Fprintf(w, " golang \n")
	})
	fmt.Println(http.ListenAndServeTLS(":8081", "server.crt", "server.key", nil))
}
