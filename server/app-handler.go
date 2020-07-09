package server

import (
	"net/http"

	"../data"
	"../service"
)

//CreateAppHandler init application handler
func CreateAppHandler(store data.Store) *AppHandler {
	return &AppHandler{store}
}

//AppHandler start api handler
type AppHandler struct {
	store data.Store
}

func (a *AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	usrServ := service.UserService{}
	router := http.NewServeMux()
	router.Handle("/api/users", &authHandler{usrServ})
	router.Handle("/api/users/login", &authHandler{usrServ})
	router.Handle("/api/user", AuthMiddleware(&userHandler{usrServ}))
	router.Handle("/api/profiles/", AuthMiddleware(&profileHandler{}))
	router.Handle("/api/articles", AuthMiddleware(&articleHandler{}))
	router.Handle("/api/tags", &tagHandler{service.TagService{a.store}})
	router.ServeHTTP(w, r)

}
