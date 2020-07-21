package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/maxim3880/real-world-api/model"
	"github.com/maxim3880/real-world-api/service"
)

type authHandler struct {
	baseHandler
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
	var got model.AuthRequestModel
	err := json.NewDecoder(r.Body).Decode(&got)
	if err != nil {
		fmt.Println(err)
		u.prepareRespData(nil, err, w)
		return
	}
	err = service.LoginValidator.Validate(got)
	if err != nil {
		fmt.Println(err)
		u.prepareRespData(nil, err, w)
		return
	}
	user, err := u.userService.LoginUser(got)
	u.prepareRespData(&user, err, w)

}

func (u *authHandler) UserRegistrationHandle(w http.ResponseWriter, r *http.Request) {
	var got model.AuthRequestModel
	err := json.NewDecoder(r.Body).Decode(&got)
	if err != nil {
		fmt.Println(err)
		u.prepareRespData(nil, err, w)
		return
	}
	err = service.RegisterValidator.Validate(got)
	if err != nil {
		fmt.Println(err)
		u.prepareRespData(nil, err, w)
		return
	}
	user, err := u.userService.RegisterUser(got)
	u.prepareRespData(&user, err, w)
}

func (u *authHandler) prepareRespData(user *model.User, err error, w http.ResponseWriter) {
	var res interface{}
	if err != nil {
		res = model.ValidationError{Body: []string{err.Error()}}
	} else {
		res = model.ResponseUserModel{User: user}
	}
	js, _ := json.Marshal(res)
	u.writeResponse(js, w, err == nil)
}

func (u *authHandler) writeResponse(data []byte, w http.ResponseWriter, success bool) {
	w.Header().Add("content-type", "application/json; charset=utf-8")
	if !success {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write(data)
}
