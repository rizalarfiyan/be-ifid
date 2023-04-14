package handler

import (
	"be-ifid/adapter"
	"be-ifid/database"
	"be-ifid/internal/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type baseHandler struct{}

func NewBaseHandler() BaseHandler {
	return &baseHandler{}
}

func (h *baseHandler) Home(ctx *fiber.Ctx) error {
	return ctx.JSON(model.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success!",
		Data: map[string]interface{}{
			"app_name": "BE IFID",
			"developer": map[string]interface{}{
				"name":      "Muhamad Rizal Arfiyan",
				"github":    "https://github.com/rizalarfiyan",
				"linkedin":  "https://linkedin.com/rizalarfiyan",
				"instagram": "https://instagram.com/rizalarfiyan",
			},
			"status": map[string]interface{}{
				"postgres": database.PostgresIsConnected(),
				"mqtt":     adapter.MQTTIsConnected(),
			},
		},
	})
}
