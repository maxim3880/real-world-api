package server

import (
	"net/http"

	"../service"
)

//CreateAppHandler init application handler
func CreateAppHandler() *AppHandler {
	return &AppHandler{}
}

//AppHandler start api handler
type AppHandler struct {
	handlers map[string]http.Handler
}

func (a *AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	usrServ := service.UserService{}
	router := http.NewServeMux()
	router.Handle("/api/users", &authHandler{usrServ})
	router.Handle("/api/users/login", &authHandler{usrServ})
	router.Handle("/api/user", AuthMiddleware(&userHandler{usrServ}))
	router.Handle("/api/profiles/", AuthMiddleware(&profileHandler{}))
	router.Handle("/api/articles", AuthMiddleware(&articleHandler{}))
	router.Handle("/api/tags", &tagHandler{})
	router.ServeHTTP(w, r)

}
