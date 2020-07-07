package server

import (
	"net/http"
)

type tagHandler struct {
}

func (u *tagHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		u.GetTags(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}

}

func (u *tagHandler) GetTags(w http.ResponseWriter, r *http.Request) {
}
