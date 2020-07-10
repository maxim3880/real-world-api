package server

import (
	"encoding/json"
	"net/http"

	"../model"
	"../service"
)

type tagHandler struct {
	service service.TagService
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
	tags := u.service.GetAll()
	respModel := model.ResponseTagModel{}
	for _, val := range tags {
		respModel.Tags = append(respModel.Tags, val.Name)
	}
	w.Header().Add("content-type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(respModel)
}
