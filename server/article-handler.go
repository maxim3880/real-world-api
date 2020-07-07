package server

import (
	"net/http"
)

type articleHandler struct {
}

func (u *articleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		u.GetArcticles(w, r)
	case http.MethodPost:
		u.CreateArticle(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}

}

func (u *articleHandler) GetArcticles(w http.ResponseWriter, r *http.Request) {
}

func (u *articleHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {

}
