package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maxim3880/real-world-api/model"
	"github.com/maxim3880/real-world-api/service"
)

type userHandler struct {
	userService service.UserService
}

func (u *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		u.GetUser(w, r)
	case http.MethodPut:
		u.UpdateUser(w, r)
	}

}

func (u *userHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	userID, ok := claim["userID"].(float64)
	if !ok {
		fmt.Println(userID)
	}
	user, err := u.userService.GetUserByID(int(userID))
	u.prepareRespData(&user, err, w)
}

func (u *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var got model.UpdateUserRequestModel
	err := json.NewDecoder(r.Body).Decode(&got)
	if err != nil {
		fmt.Println(err)
		u.prepareRespData(nil, err, w)
		return
	}
	userID, _ := claim["userID"].(float64)
	user, err := u.userService.UpdateUser(got, int(userID))
	u.prepareRespData(&user, err, w)
}

func (u *userHandler) prepareRespData(user *model.User, err error, w http.ResponseWriter) {
	var res interface{}
	if err != nil {
		res = model.ValidationError{Body: []string{err.Error()}}
	} else {
		user.Token = service.GenerateJwtTocken(user.Email, user.ID)
		res = model.ResponseUserModel{User: user}
	}
	js, _ := json.Marshal(res)
	u.writeResponse(js, w, err == nil)
}

func (u *userHandler) writeResponse(data []byte, w http.ResponseWriter, success bool) {
	w.Header().Add("content-type", "application/json; charset=utf-8")
	if !success {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write(data)
}
