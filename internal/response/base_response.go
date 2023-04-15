package response

import (
	"be-ifid/utils"
	"net/http"
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

func NewBindingError() *BaseResponse {
	return &BaseResponse{
		Code:    http.StatusBadRequest,
		Message: "Binding Validation Error",
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
