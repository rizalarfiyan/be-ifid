package handler

import (
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
		Data:    nil,
	})
}
