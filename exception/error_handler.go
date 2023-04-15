package exception

import (
	"be-ifid/internal/response"
	"errors"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	var data interface{}
	if strings.EqualFold(err.Error(), http.StatusText(code)) {
		data = nil
	} else {
		data = err.Error()
	}

	return ctx.JSON(response.BaseResponse{
		Code:    code,
		Message: http.StatusText(code),
		Data:    data,
	})
}
