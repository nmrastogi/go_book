package main

import (
	"log"

	"github.com/nmrastogi/go_book/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
