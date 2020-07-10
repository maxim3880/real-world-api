package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"../data"
	"../model"
	"github.com/stretchr/testify/assert"
)

func TestTagHandler(t *testing.T) {

	appHandler := CreateAppHandler(data.CreateImMemmoryStore())
	t.Run("get tags list", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/api/tags", nil)

		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)

		var result model.ResponseTagModel
		json.NewDecoder(response.Body).Decode(&result)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, len(result.Tags), 2)
		assert.Contains(t, result.Tags, "angularjs")
		assert.Contains(t, result.Tags, "reactjs")
	})
	t.Run("get tags list", func(t *testing.T) {
		reqBody := `{}`
		req, _ := http.NewRequest(http.MethodPost, "/api/tags", strings.NewReader(reqBody))

		response := httptest.NewRecorder()

		appHandler.ServeHTTP(response, req)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

}
