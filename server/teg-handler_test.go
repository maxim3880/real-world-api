package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagHandler(t *testing.T) {

	appHandler := CreateAppHandler()
	t.Run("get tags list", func(t *testing.T) {
		reqBody := `{}`
		req, _ := http.NewRequest(http.MethodGet, "/api/tags", strings.NewReader(reqBody))

		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)

		assert.Equal(t, http.StatusOK, response.Code)
	})
	t.Run("get tags list", func(t *testing.T) {
		reqBody := `{}`
		req, _ := http.NewRequest(http.MethodPost, "/api/tags", strings.NewReader(reqBody))

		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

}
