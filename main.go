package main

import (
	"log"
	"net/http"

	"./server"
)

func main() {

	if err := http.ListenAndServe(":3000", server.CreateAppHandler()); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
