package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/maxim3880/real-world-api/handler"
	"github.com/maxim3880/real-world-api/restapi"

	"github.com/maxim3880/real-world-api/server"
)

func main() {
	swagger, err := restapi.GetSwagger()
	if err != nil {
		log.Fatal("Can not get swagger spec for api")
	}
	h := handler.NewHandler()
	s := server.NewAPIServer(h, swagger)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGHUP,
		syscall.SIGQUIT)

	select {
	case err := <-s.ListenAndServe():
		log.Fatalf("Error while listen port %v", err)
	case sig := <-sigCh:
		log.Printf("terminating got [%v] signal", sig)
	}
}
