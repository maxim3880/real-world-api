package main

import (
	"log"
	"net/http"

	"github.com/maxim3880/real-world-api/data"
	"github.com/maxim3880/real-world-api/server"
)

func main() {

	store := data.CreatePostgresDbStore()
	if err := http.ListenAndServe(":3000", server.CreateAppHandler(store)); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
