package server

import (
	"net/http"
	"strings"

	"../service"
)

//CreateAppHandler init application handler
func CreateAppHandler() *AppHandler {
	return &AppHandler{
		handlers: map[string]http.Handler{
			"users": &userHandler{
				userService: service.UserService{},
			},
		},
	}
}

//AppHandler start api handler
type AppHandler struct {
	handlers map[string]http.Handler
}

func (a *AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handlerPth := strings.Split(r.URL.Path[1:], "/")
	if len(handlerPth) > 1 {
		if handler, ok := a.handlers[handlerPth[1]]; ok {
			handler.ServeHTTP(w, r)
			return
		}
	}

	http.Error(w, "Not Found", http.StatusNotFound)

}
