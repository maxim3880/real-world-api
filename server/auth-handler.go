package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"../model"
	"../service"
)

type authHandler struct {
	userService service.UserService
}

func (u *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	default:
		{
			w.WriteHeader(http.StatusNotFound)
		}
	}

}

func (u *authHandler) UserLoginHandle(w http.ResponseWriter, r *http.Request) {
	var got model.RequestUserModel
	err := json.NewDecoder(r.Body).Decode(&got)
	if err != nil {

	}
	u.userService.LoginUser(got.User)
	got.User.Token = service.GenerateJwtTocken(got.User.Email)
	w.Header().Add("content-type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(got)

}

func (u *authHandler) UserRegistrationHandle(w http.ResponseWriter, r *http.Request) {
	var got model.RequestUserModel
	err := json.NewDecoder(r.Body).Decode(&got)
	if err != nil {

	}
	u.userService.LoginUser(got.User)
	got.User.Token = service.GenerateJwtTocken(got.User.Email)
	w.Header().Set("content-type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(got)
}
