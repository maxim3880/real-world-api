package server

import (
	"errors"
	"net/http"
	"strings"

	"github.com/maxim3880/real-world-api/model"
	"github.com/maxim3880/real-world-api/service"
)

type profileHandler struct {
	baseHandler
	service service.ProfileService
}

func (u *profileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		u.GetUserProfile(w, r)
	case http.MethodPost:
		u.CreateFollowUser(w, r)
	case http.MethodDelete:
		u.DeleteFollowUser(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}

}

func (u *profileHandler) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	pth := strings.Split(r.URL.Path[1:], "api/profiles/")
	prof, err := u.service.GetProfileByUsername(pth[1])
	if err == nil {
		u.prepareRespData(model.ProfileResponse{Profile: prof}, err, w)
		return
	}
	u.prepareRespData(prof, err, w)
}

func (u *profileHandler) CreateFollowUser(w http.ResponseWriter, r *http.Request) {
	pth := strings.Split(r.URL.Path[1:], "api/profiles/")
	userID, ok := claim["userID"].(float64)
	if !ok {
		u.prepareRespData(nil, errors.New("user id in token not exists"), w)
		return
	}
	val := strings.Split(pth[1], "/")
	prof, err := u.service.CreateProfileFollow(val[0], int(userID))
	if err == nil {
		u.prepareRespData(model.ProfileResponse{Profile: prof}, err, w)
		return
	}
	u.prepareRespData(prof, err, w)
}

func (u *profileHandler) DeleteFollowUser(w http.ResponseWriter, r *http.Request) {
	pth := strings.Split(r.URL.Path[1:], "api/profiles/")
	val := strings.Split(pth[1], "/")
	userID, ok := claim["userID"].(float64)
	if !ok {
		u.prepareRespData(nil, errors.New("user id in token not exists"), w)
		return
	}
	prof, err := u.service.DeleteProfileFollow(val[0], int(userID))
	if err == nil {
		u.prepareRespData(model.ProfileResponse{Profile: prof}, err, w)
		return
	}
	u.prepareRespData(prof, err, w)
}
