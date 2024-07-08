package main

import (
	"database/sql"
	"github.com/lib/pq"
	"log"
)

func main() {
	db, err := sql.Open("postgre", "user=pggotest dbname=pqgotest sslmode=verify=full")
	if err != nil {
		log.Fatal(err)
	}
	age := 21
	db.Query("SELECT name FROM users WHERE age = $1", age)
	// ...
}

func init() {
	sql.Register("postgres", &pq.Driver{})
}
