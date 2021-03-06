package server

import (
	"net/http"

	"github.com/maxim3880/real-world-api/data"
	"github.com/maxim3880/real-world-api/service"
)

//CreateAppHandler init application handler
func CreateAppHandler(store data.Store) *AppHandler {
	service.SetValidatorTags()
	return &AppHandler{store}
}

//AppHandler start api handler
type AppHandler struct {
	store data.Store
}

func (a *AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	router := http.NewServeMux()
	for p, hand := range a.GetRoutesWithHanders() {
		router.Handle(p, hand)
	}
	router.ServeHTTP(w, r)

}

//GetRoutesWithHanders return all api path with endpoint handler
func (a *AppHandler) GetRoutesWithHanders() map[string]http.Handler {
	baseHandler := baseHandler{}
	userserv := service.UserService{Store: a.store}
	return map[string]http.Handler{
		"/api/users":       validationMiddleware(&authHandler{baseHandler, userserv}),
		"/api/users/login": validationMiddleware(&authHandler{baseHandler, userserv}),
		"/api/user":        validationMiddleware(&userHandler{baseHandler, userserv}),
		"/api/profiles/":   validationMiddleware(&profileHandler{baseHandler: baseHandler, service: service.ProfileService{Store: a.store}}),
		"/api/articles":    validationMiddleware(&articleHandler{baseHandler, service.ArticleService{Store: a.store}}),
		"/api/tags":        &tagHandler{service.TagService{Store: a.store}},
	}
}
