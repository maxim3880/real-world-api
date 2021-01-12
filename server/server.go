package server

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gofiber/fiber/v2"
	"github.com/maxim3880/real-world-api/middleware"
	"github.com/maxim3880/real-world-api/restapi"
)

type server struct {
	http *fiber.App
}

func NewAPIServer(h restapi.ServerInterface, swagger *openapi3.Swagger) *server {

	app := fiber.New()

	swagger.Servers = nil
	app.Use(middleware.OapiRequestValidator(swagger))

	restapi.RegisterHandlers(app, h)
	return &server{http: app}
}

func (s *server) ListenAndServe() chan error {
	errCh := make(chan error)

	go func() {
		listenAddress := fmt.Sprintf("%s:%s", "", "8080")
		errCh <- s.http.Listen(listenAddress)
	}()

	return errCh
}
