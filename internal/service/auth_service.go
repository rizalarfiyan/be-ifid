package service

import (
	"be-ifid/config"
	"be-ifid/internal/repository"
	"context"
)

type authService struct {
	ctx  context.Context
	conf *config.Config
	repo repository.AuthRepository
}

func NewAuthService(ctx context.Context, conf *config.Config, repo repository.AuthRepository) AuthService {
	return &authService{
		ctx,
		conf,
		repo,
	}
}
