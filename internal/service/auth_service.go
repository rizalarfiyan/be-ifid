package service

import "be-ifid/internal/request"

type AuthService interface {
	Login(req request.AuthRequest) error
}
