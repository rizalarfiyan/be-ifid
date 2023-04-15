package handler

import (
	"be-ifid/config"
	"be-ifid/database"
	"be-ifid/internal/response"
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type authHandler struct {
	ctx      context.Context
	conf     *config.Config
	postgres *sqlx.DB
	redis    database.RedisInstance
}

func NewAuthHandler(ctx context.Context, conf *config.Config, postgres *sqlx.DB, redis database.RedisInstance) AuthHandler {
	return &authHandler{
		ctx,
		conf,
		postgres,
		redis,
	}
}

func (h *authHandler) Login(ctx *fiber.Ctx) error {
	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Login!",
		Data:    nil,
	})
}
