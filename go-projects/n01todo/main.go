package main

func main() {
	type TODOs struct {
		Id          int    `json:"id"`
		Description string `json:"description"`
		Status      bool   `json:"status"`
	}
}
