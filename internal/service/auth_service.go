package service

import (
	"be-ifid/internal/request"
	"be-ifid/internal/response"
)

type AuthService interface {
	Login(req request.AuthRequest) error
	Callback(token string) (*response.AuthCallbackResponse, error)
}
