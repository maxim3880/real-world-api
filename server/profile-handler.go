package server

import (
	"net/http"
)

type profileHandler struct {
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
}

func (u *profileHandler) CreateFollowUser(w http.ResponseWriter, r *http.Request) {

}

func (u *profileHandler) DeleteFollowUser(w http.ResponseWriter, r *http.Request) {

}
