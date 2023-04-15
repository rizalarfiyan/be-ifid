package handler

import (
	"be-ifid/config"
	"be-ifid/constant"
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
		service.NewAuthService(ctx, conf, repo, redis),
	}
}

func (h *authHandler) Login(ctx *fiber.Ctx) error {
	var req request.AuthRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return response.NewBindingError()
	}

	err = req.Validate()
	if err != nil {
		return response.NewValidationError(err)
	}

	err = h.service.Login(req)
	if err != nil {
		return err
	}

	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Please check your email to verify your account",
		Data:    nil,
	})
}

func (h *authHandler) Callback(ctx *fiber.Ctx) error {
	token := ctx.Query("token", "")
	if token == "" {
		return response.NewErrorMessage(http.StatusBadRequest, "Token is required!", nil)
	}

	if len(token) != constant.AuthKeyLength {
		return response.NewErrorMessage(http.StatusUnprocessableEntity, "Token invalid!", nil)
	}

	err := h.service.Callback(token)
	if err != nil {
		return err
	}

	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Login!",
		Data:    nil,
	})
}
