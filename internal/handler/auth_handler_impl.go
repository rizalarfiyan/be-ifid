package handler

import (
	"be-ifid/config"
	"be-ifid/database"
	"be-ifid/internal/repository"
	"be-ifid/internal/request"
	"be-ifid/internal/response"
	"be-ifid/internal/service"
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type authHandler struct {
	conf    *config.Config
	service service.AuthService
}

func NewAuthHandler(ctx context.Context, conf *config.Config, postgres *sqlx.DB, redis database.RedisInstance) AuthHandler {
	repo := repository.NewAuthRepository(ctx, conf, postgres, redis)
	return &authHandler{
		conf,
		service.NewAuthService(ctx, conf, repo),
	}
}

func (h *authHandler) Login(ctx *fiber.Ctx) error {
	var request request.AuthRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		return response.NewBindingError()
	}

	err = request.Validate()
	if err != nil {
		return response.NewValidationError(err)
	}

	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Login!",
		Data:    nil,
	})
}
