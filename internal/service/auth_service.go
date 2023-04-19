package service

import (
	"be-ifid/internal/model"
	"be-ifid/internal/request"
	"be-ifid/internal/response"
)

type AuthService interface {
	Login(req request.LoginRequest) error
	Callback(token string) (*response.AuthTokenResponse, error)
	FirstUser(user model.JWTAuthPayload, req request.FirstUserRequest) (*response.AuthTokenResponse, error)
	Me(user model.JWTAuthPayload) (*response.AuthMeResponse, error)
}
