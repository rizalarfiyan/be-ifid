package exception

import (
	"be-ifid/internal/response"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	return ctx.JSON(response.BaseResponse{
		Code:    code,
		Message: http.StatusText(code),
		Data:    err.Error(),
	})
}
