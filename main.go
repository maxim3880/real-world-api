package main

import (
	"log"
	"net/http"

	"./server"
)

func main() {

	//Jwt middleware for check auth tocken
	//http.Handle("/", server.AuthMiddleware(server.CreateAppHandler()))

	if err := http.ListenAndServe(":3000", server.CreateAppHandler()); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
