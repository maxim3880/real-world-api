package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"../model"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler(t *testing.T) {

	appHandler := CreateAppHandler()
	t.Run("check POST user login", func(t *testing.T) {
		reqBody := `{
			"user":{
			  "email": "jake@jake.jake",
			  "password": "jakejake"
			}
		  }`
		req, _ := http.NewRequest(http.MethodPost, "/api/users/login", strings.NewReader(reqBody))
		req.Header.Set("content-type", "application/json; charset=utf-8")
		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)
		var result model.ResponseUserModel
		json.NewDecoder(response.Body).Decode(&result)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "jake@jake.jake", result.User.Email)
		assert.NotEmpty(t, result.User.Token)
	})

	t.Run("check POST user registration", func(t *testing.T) {
		reqBody := `{
			"user":{
			  "username": "Jacob",
			  "email": "jake@jake.jake",
			  "password": "jakejake"
			}
		  }`
		req, _ := http.NewRequest(http.MethodPost, "/api/users", strings.NewReader(reqBody))
		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)
		var result model.ResponseUserModel
		json.NewDecoder(response.Body).Decode(&result)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "jake@jake.jake", result.User.Email)
		assert.NotEmpty(t, result.User.Token)
	})

	t.Run("check PUT user data", func(t *testing.T) {
		reqBody := `{
			"user":{
			}
		  }`
		req, _ := http.NewRequest(http.MethodPut, "/api/users", strings.NewReader(reqBody))
		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)

		assert.Equal(t, http.StatusUnauthorized, response.Code)
	})
}
