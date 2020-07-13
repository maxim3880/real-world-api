package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"../data"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler(t *testing.T) {

	appHandler := CreateAppHandler(data.CreateImMemmoryStore("TestUserHandlerDataSource"))
	t.Run("check PUT user data", func(t *testing.T) {
		reqBody := `{}`
		req, _ := http.NewRequest(http.MethodPut, "/api/user", strings.NewReader(reqBody))
		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)

		assert.Equal(t, http.StatusUnauthorized, response.Code)
	})
}
