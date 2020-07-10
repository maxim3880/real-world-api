package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"../data"
	"github.com/stretchr/testify/assert"
)

func TestAppHandler(t *testing.T) {
	appHandler := CreateAppHandler(data.CreatePostgresDbStore())
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
