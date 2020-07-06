package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"../model"
	"../service"
)

type userHandler struct {
	userService service.UserService
}

func (u *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handlerPth := strings.Split(r.URL.Path[1:], "/")

	switch r.Method {
	case http.MethodPost:
		{
			if len(handlerPth) > 2 && handlerPth[2] == "login" {
				u.UserLoginHandle(w, r)
			} else {
				u.UserRegistrationHandle(w, r)
			}
		}
	case http.MethodPut:
		u.UpdateUser(w, r)
	}

}

func (u *userHandler) UserLoginHandle(w http.ResponseWriter, r *http.Request) {
	var got model.RequestUserModel
	if err := json.NewDecoder(r.Body).Decode(&got); err == nil {
		u.userService.LoginUser(got.User)
		got.User.Token = service.GenerateJwtTocken(got.User.Email)
		w.Header().Add("content-type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(got)
	}

}

func (u *userHandler) UserRegistrationHandle(w http.ResponseWriter, r *http.Request) {
	var got model.RequestUserModel
	if err := json.NewDecoder(r.Body).Decode(&got); err == nil {
		u.userService.LoginUser(got.User)
		w.Header().Set("content-type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(got)
	}
}

func (u *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

}
