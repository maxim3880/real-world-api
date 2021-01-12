package handler

import (
	"github.com/gofiber/fiber/v2"
)

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

// Get current user
// (GET /users)
func (h *handler) GetUsers(ctx *fiber.Ctx) error {
	return nil
}
