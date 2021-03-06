package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/maxim3880/real-world-api/model"
	"github.com/maxim3880/real-world-api/service"
)

var claim jwt.MapClaims = nil

func validationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			e := recover()
			if e == nil {
				next.ServeHTTP(w, r)
			} else {
				if errors.Is(e.(error), &model.UnauthorizedError{}) {
					writeUnauthorized(e, w)
				} else {
					writeBadRequest(e, w)
				}
			}
		}()
		for _, validators := range getValidationChain() {
			validators(w, r)
		}

	})
}

func getValidationChain() []http.HandlerFunc {
	return []http.HandlerFunc{
		validateEmptyBody,
		validateAuthToken,
	}
}

func validateEmptyBody(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Body == http.NoBody && !contains(getRoutesWithoutBody(), r.URL.Path) {
		err := model.ValidationError{Body: []string{}}
		err.Body = append(err.Body, model.ErrorMsgs["bodyNotEmpty"])
		panic(err)
	}
}

func validateAuthToken(w http.ResponseWriter, r *http.Request) {
	if contains(getRoutesWithAuth(), r.URL.Path) {
		resErr := &model.UnauthorizedError{}
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			resErr.Message = "Authorization header required"
			panic(resErr)
		}
		headerVal := strings.Fields(authHeader)
		if len(headerVal) < 2 {
			resErr.Message = "Authorization header not contain required value"
			panic(resErr)
		} else if strings.ToLower(headerVal[0]) != "token" {
			resErr.Message = "Authorization not start from 'Token' word"
			panic(resErr)
		}
		token := headerVal[1]
		parsedToken, err := jwt.Parse(token, service.GetJwtValidationKey)
		if err != nil {
			resErr.Message = err.Error()
			panic(resErr)
		}
		claim = parsedToken.Claims.(jwt.MapClaims)
	}
}

func writeBadRequest(err interface{}, w http.ResponseWriter) {
	w.Header().Add("content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(err)
}

func writeUnauthorized(err interface{}, w http.ResponseWriter) {
	text := err.(error).Error()
	res := model.ValidationError{Header: []string{text}}
	w.Header().Add("content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(res)
}

func getRoutesWithAuth() map[string]int {
	return map[string]int{
		"/api/user":      0,
		"/api/profiles/": 1,
	}
}

func getRoutesWithoutBody() map[string]int {
	return map[string]int{
		"/api/profiles/": 1,
	}
}

func contains(s map[string]int, e string) bool {
	for pth, comType := range s {
		switch comType {
		case 0:
			if pth == e {
				return true
			}
		case 1:
			if strings.HasPrefix(e, pth) {
				return true
			}
		}
	}
	return false
}
