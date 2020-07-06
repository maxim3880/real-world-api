package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppHandler(t *testing.T) {
	t.Run("check is handler by path correct", func(t *testing.T) {
		userHandlerMock := getHandlerMockStub()

		errHandlerMock := getHandlerMockStub()

		appHandler := &AppHandler{map[string]http.Handler{
			"users":      userHandlerMock,
			"error_path": errHandlerMock,
		}}
		req, _ := http.NewRequest(http.MethodGet, "/api/users", nil)
		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)

		userHandlerMock.AssertNumberOfCalls(t, "ServeHTTP", 1)
		errHandlerMock.AssertNumberOfCalls(t, "ServeHTTP", 0)
	})
	t.Run("check for 404 when path not correct", func(t *testing.T) {
		userHandlerMock := getHandlerMockStub()

		appHandler := &AppHandler{map[string]http.Handler{
			"users": userHandlerMock,
		}}
		req, _ := http.NewRequest(http.MethodGet, "/api/not-user", nil)
		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)

		assert.Equal(t, http.StatusNotFound, response.Code)
		userHandlerMock.AssertNumberOfCalls(t, "ServeHTTP", 0)
	})
	t.Run("check 404 for empty route path", func(t *testing.T) {
		userHandlerMock := getHandlerMockStub()

		appHandler := &AppHandler{map[string]http.Handler{
			"users": userHandlerMock,
		}}
		req, _ := http.NewRequest(http.MethodGet, "/api", nil)
		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)

		assert.Equal(t, http.StatusNotFound, response.Code)
		userHandlerMock.AssertNumberOfCalls(t, "ServeHTTP", 0)
	})
}
