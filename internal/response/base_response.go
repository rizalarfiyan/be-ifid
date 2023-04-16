package response

import (
	"be-ifid/utils"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewError(code int, data interface{}) *BaseResponse {
	return &BaseResponse{
		Code: code,
		Data: data,
	}
}

func NewErrorMessage(code int, message string, data interface{}) *BaseResponse {
	return &BaseResponse{
		code,
		message,
		data,
	}
}

func NewBindingError(err error) *BaseResponse {
	code := http.StatusBadRequest
	message := "Binding Validation Error"

	var fiberError *fiber.Error
	if errors.As(err, &fiberError) {
		code = fiberError.Code
		message = fiberError.Message
	}

	return &BaseResponse{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

func NewValidationError(err error) *BaseResponse {
	return &BaseResponse{
		Code:    http.StatusBadRequest,
		Message: "Validation Error",
		Data:    utils.ParseValidation(err),
	}
}

func (res *BaseResponse) Error() string {
	return res.Message
}
