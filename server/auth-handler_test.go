package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/maxim3880/real-world-api/data"
	"github.com/maxim3880/real-world-api/model"
	"github.com/stretchr/testify/assert"
)

func TestAuthLoginHandler(t *testing.T) {

	appHandler := CreateAppHandler(data.CreateImMemmoryStore("TestAuthLoginHandlerDataSource"))
	t.Run("check POST user login", func(t *testing.T) {
		reqBody := `{
			"user":{
			  "email": "maxim3880@gmail.com",
			  "password": "1234"
			}
		  }`
		req, _ := http.NewRequest(http.MethodPost, "/api/users/login", strings.NewReader(reqBody))
		req.Header.Set("content-type", "application/json; charset=utf-8")
		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)
		var result model.ResponseUserModel
		json.NewDecoder(response.Body).Decode(&result)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "maxim3880@gmail.com", result.User.Email)
		assert.NotEmpty(t, result.User.Token)
	})
	t.Run("check POST user login", func(t *testing.T) {
		reqBody := `{
			"user":{
			  "email": "notexistsuser@gmail.com",
			  "password": "12345"
			}
		  }`
		req, _ := http.NewRequest(http.MethodPost, "/api/users/login", strings.NewReader(reqBody))
		req.Header.Set("content-type", "application/json; charset=utf-8")
		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)
		var result model.ValidationError
		json.NewDecoder(response.Body).Decode(&result)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.NotEmpty(t, result.Body)
		assert.Equal(t, model.ErrorMsgs["userNotExists"], result.Body[0])
	})
	t.Run("check POST user login with empty password", func(t *testing.T) {
		reqBody := `{
			"user":{
			  "email": "maxim3880@gmail.com",
			  "password": ""
			}
		  }`
		req, _ := http.NewRequest(http.MethodPost, "/api/users/login", strings.NewReader(reqBody))
		req.Header.Set("content-type", "application/json; charset=utf-8")
		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)
		var result model.ValidationError
		json.NewDecoder(response.Body).Decode(&result)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.NotEmpty(t, result.Body)
		assert.Equal(t, model.ErrorMsgs["emptyPassword"], result.Body[0])
	})
	t.Run("check POST user login with empty email", func(t *testing.T) {
		reqBody := `{
			"user":{
			  "email": "",
			  "password": "1234"
			}
		  }`
		req, _ := http.NewRequest(http.MethodPost, "/api/users/login", strings.NewReader(reqBody))
		req.Header.Set("content-type", "application/json; charset=utf-8")
		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)
		var result model.ValidationError
		json.NewDecoder(response.Body).Decode(&result)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.NotEmpty(t, result.Body)
		assert.Equal(t, model.ErrorMsgs["emptyEmail"], result.Body[0])
	})

}

func TestAuthSignupHandler(t *testing.T) {
	appHandler := CreateAppHandler(data.CreateImMemmoryStore("TestAuthSignupHandlerDataSource"))
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
	t.Run("check POST user registration", func(t *testing.T) {
		reqBody := `{
			"user":{
			  "username": "",
			  "email": "jake@jake.jake",
			  "password": "jakejake"
			}
		  }`
		req, _ := http.NewRequest(http.MethodPost, "/api/users", strings.NewReader(reqBody))
		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)
		var result model.ValidationError
		json.NewDecoder(response.Body).Decode(&result)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.NotEmpty(t, result.Body)
		assert.Equal(t, model.ErrorMsgs["emptyUsername"], result.Body[0])
	})
	t.Run("check POST user registration", func(t *testing.T) {
		reqBody := ``
		req, _ := http.NewRequest(http.MethodPost, "/api/users", strings.NewReader(reqBody))
		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)
		var result model.ValidationError
		json.NewDecoder(response.Body).Decode(&result)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.NotEmpty(t, result.Body)
		assert.Equal(t, model.ErrorMsgs["bodyNotEmpty"], result.Body[0])
	})
}
