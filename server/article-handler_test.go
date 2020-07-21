package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/maxim3880/real-world-api/data"
	"github.com/maxim3880/real-world-api/model"
	"github.com/maxim3880/real-world-api/service"
	"github.com/stretchr/testify/assert"
)

func TestArticleHandler(t *testing.T) {

	store := data.CreateImMemmoryStore("TestArticleHandlerDataSource")
	tkn := "Token " + service.GenerateJwtTocken("test-email@test.com", 1)
	appHandler := CreateAppHandler(store)
	t.Run("get articles data without filter and token", func(t *testing.T) {
		//mockStore := new(data.MockDBStore)
		req, _ := http.NewRequest(http.MethodGet, "/api/articles", nil)
		response := httptest.NewRecorder()

		appHandler := CreateAppHandler(store)
		appHandler.ServeHTTP(response, req)

		var result model.MultiArticleResponse
		json.NewDecoder(response.Body).Decode(&result)

		assert.Equal(t, http.StatusOK, response.Code)
		//assert.Equal(t, 2, result.ArticlesCount)
		//assert.GreaterOrEqual(t, result.Articles[0], result.Articles[1])
	})
	t.Run("get articles data with token and without filter", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/api/articles", nil)
		req.Header.Add("Authorization", tkn)
		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)

		var result model.MultiArticleResponse
		json.NewDecoder(response.Body).Decode(&result)

		assert.Equal(t, http.StatusOK, response.Code)
		//assert.Equal(t, 2, result.ArticlesCount)
		//assert.GreaterOrEqual(t, result.Articles[0], result.Articles[1])
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
