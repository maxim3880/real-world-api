package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"../data"
	"../service"
	"github.com/stretchr/testify/assert"
)

func TestArticleHandler(t *testing.T) {

	tkn := "Bearer " + service.GenerateJwtTocken("test-email@test.com", 1)
	appHandler := CreateAppHandler(data.CreateImMemmoryStore("TestArticleHandlerDataSource"))
	t.Run("get profile data", func(t *testing.T) {
		reqBody := `{}`
		req, _ := http.NewRequest(http.MethodGet, "/api/articles", strings.NewReader(reqBody))

		req.Header.Add("Authorization", tkn)
		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)

		assert.Equal(t, http.StatusOK, response.Code)
	})
	t.Run("post profile follow data", func(t *testing.T) {
		reqBody := `{}`
		req, _ := http.NewRequest(http.MethodPost, "/api/articles", strings.NewReader(reqBody))
		req.Header.Add("Authorization", tkn)

		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)

		assert.Equal(t, http.StatusOK, response.Code)
	})
	t.Run("post profile follow data", func(t *testing.T) {
		reqBody := `{}`
		req, _ := http.NewRequest(http.MethodPut, "/api/articles", strings.NewReader(reqBody))
		req.Header.Add("Authorization", tkn)

		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
	t.Run("post profile follow data", func(t *testing.T) {
		reqBody := `{}`
		req, _ := http.NewRequest(http.MethodDelete, "/api/articles", strings.NewReader(reqBody))
		req.Header.Add("Authorization", tkn)

		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
}
