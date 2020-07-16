package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/maxim3880/real-world-api/data"
	"github.com/maxim3880/real-world-api/service"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler(t *testing.T) {

	tkn := "Token " + service.GenerateJwtTocken("test-email@test.com", 1)
	appHandler := CreateAppHandler(data.CreateImMemmoryStore("TestUserHandlerDataSource"))
	t.Run("check GET user data", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/api/user", nil)
		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)

		assert.Equal(t, http.StatusUnauthorized, response.Code)
	})
	t.Run("check GET user data", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/api/user", nil)
		req.Header.Add("Authorization", tkn)
		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)

		assert.Equal(t, http.StatusOK, response.Code)
	})
	t.Run("check PUT user data", func(t *testing.T) {
		reqBody := `{}`
		req, _ := http.NewRequest(http.MethodPut, "/api/user", strings.NewReader(reqBody))
		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)

		assert.Equal(t, http.StatusUnauthorized, response.Code)
	})
}
