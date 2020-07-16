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

func TestAppHandler(t *testing.T) {
	appHandler := CreateAppHandler(data.CreateImMemmoryStore("TestAppHandlerDataSource"))
	t.Run("check POST send empty body", func(t *testing.T) {
		reqBody := ``
		req, _ := http.NewRequest(http.MethodPost, "/api/users/login", strings.NewReader(reqBody))
		req.Header.Set("content-type", "application/json; charset=utf-8")
		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)
		var result model.ValidationError
		json.NewDecoder(response.Body).Decode(&result)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.NotEmpty(t, result.Body)
		assert.Equal(t, model.ErrorMsgs["bodyNotEmpty"], result.Body[0])
	})
	t.Run("check is handler by path correct", func(t *testing.T) {

		req, _ := http.NewRequest(http.MethodGet, "/api", nil)
		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
	t.Run("check for 404 when path not correct", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/api/not-user", nil)
		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
	t.Run("check 404 for empty route path", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
}
