package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"../service"
	"github.com/stretchr/testify/assert"
)

func TestProfileHandler(t *testing.T) {

	tkn := "Bearer " + service.GenerateJwtTocken("test-email@test.com")
	appHandler := CreateAppHandler()
	t.Run("get profile data", func(t *testing.T) {
		reqBody := `{}`
		req, _ := http.NewRequest(http.MethodGet, "/api/profiles/testname", strings.NewReader(reqBody))

		req.Header.Add("Authorization", tkn)
		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)

		assert.Equal(t, http.StatusOK, response.Code)
	})
	t.Run("post profile follow data", func(t *testing.T) {
		reqBody := `{}`
		req, _ := http.NewRequest(http.MethodPost, "/api/profiles/testname/follow", strings.NewReader(reqBody))
		req.Header.Add("Authorization", tkn)

		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)

		assert.Equal(t, http.StatusOK, response.Code)
	})
	t.Run("delete profile follow data", func(t *testing.T) {
		reqBody := `{}`
		req, _ := http.NewRequest(http.MethodDelete, "/api/profiles/testname/follow", strings.NewReader(reqBody))
		req.Header.Add("Authorization", tkn)

		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)

		assert.Equal(t, http.StatusOK, response.Code)
	})
	t.Run("delete profile follow data", func(t *testing.T) {
		reqBody := `{}`
		req, _ := http.NewRequest(http.MethodPut, "/api/profiles/testname/follow", strings.NewReader(reqBody))
		req.Header.Add("Authorization", tkn)

		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
}
