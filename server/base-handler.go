package server

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/maxim3880/real-world-api/model"
)

//baseHandler represent base class of all handlers
type baseHandler struct {
	claim jwt.MapClaims
}

func (b *baseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func (b *baseHandler) prepareRespData(data interface{}, err error, w http.ResponseWriter) {
	var res interface{}
	if err != nil {
		res = model.ValidationError{Body: []string{err.Error()}}
	} else {
		res = data
	}
	js, _ := json.Marshal(res)
	b.writeResponse(js, w, err == nil)
}

func (b *baseHandler) writeResponse(data []byte, w http.ResponseWriter, success bool) {
	w.Header().Add("content-type", "application/json; charset=utf-8")
	if !success {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write(data)
}
