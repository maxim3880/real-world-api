package server

import (
	"net/http"

	"github.com/maxim3880/real-world-api/model"
	"github.com/maxim3880/real-world-api/service"
)

type articleHandler struct {
	baseHandler
	service service.ArticleService
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
	res := u.service.GetAll()
	u.prepareRespData(model.MultiArticleResponse{Articles: res, ArticlesCount: len(res)}, nil, w)

}

func (u *articleHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {

}
