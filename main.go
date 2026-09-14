package main

import (
	"log"
	"net/http"

	"github.com/goobnoobler/server/server"
)

func main() {
	r := server.Request()
	log.Fatal(http.ListenAndServe(":8080", r))
}
