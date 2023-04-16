package service

import (
	"be-ifid/internal/request"
	"be-ifid/internal/response"
)

type AuthService interface {
	Login(req request.LoginRequest) error
	Callback(token string) (*response.AuthCallbackResponse, error)
}
