package server

import (
	"fmt"
	"net/http"

	"../service"
)

type userHandler struct {
	userService service.UserService
}

func (u *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(claim)
	switch r.Method {
	case http.MethodGet:
		u.UpdateUser(w, r)
	case http.MethodPut:
		u.UpdateUser(w, r)
	}

}

func (u *userHandler) GetUser(w http.ResponseWriter, r *http.Request) {
}

func (u *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

}
