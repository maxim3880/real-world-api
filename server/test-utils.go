package server

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type handlerMock struct {
	mock.Mock
}

func (h *handlerMock) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Called()
}

func getHandlerMockStub() (mockObj *handlerMock) {
	mockObj = new(handlerMock)
	mockObj.On("ServeHTTP", mock.Anything)
	return
}
