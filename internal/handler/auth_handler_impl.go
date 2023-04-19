package handler

import (
	"be-ifid/config"
	"be-ifid/constant"
	"be-ifid/database"
	"be-ifid/internal/model"
	"be-ifid/internal/repository"
	"be-ifid/internal/request"
	"be-ifid/internal/response"
	"be-ifid/internal/service"
	"context"
	"net/http"
	"regexp"

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
	var req request.LoginRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return response.NewBindingError(err)
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

	match, _ := regexp.MatchString(`^[a-zA-Z0-9-_]+$`, token)
	if len(token) != constant.AuthKeyLength || !match {
		return response.NewErrorMessage(http.StatusUnprocessableEntity, "Token invalid!", nil)
	}

	data, err := h.service.Callback(token)
	if err != nil {
		return err
	}

	message := "Welcome to IFID!"
	if data.IsNew {
		message = "Welcome to IFID! Get started by adding your first account."
	}
	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	})
}

func (h *authHandler) FirstUser(ctx *fiber.Ctx) error {
	var user model.JWTAuthPayload
	err := user.GetFromFiber(ctx)
	if err != nil {
		return err
	}

	if !user.IsNew {
		return response.NewErrorMessage(http.StatusUnprocessableEntity, "User already registered", nil)
	}

	var req request.FirstUserRequest
	err = ctx.BodyParser(&req)
	if err != nil {
		return response.NewBindingError(err)
	}

	err = req.Validate()
	if err != nil {
		return response.NewValidationError(err)
	}

	data, err := h.service.FirstUser(user, req)
	if err != nil {
		return err
	}

	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusCreated,
		Message: "Success save data to database",
		Data:    data,
	})
}

func (h *authHandler) Me(ctx *fiber.Ctx) error {
	var user model.JWTAuthPayload
	err := user.GetFromFiber(ctx)
	if err != nil {
		return err
	}

	data, err := h.service.Me(user)
	if err != nil {
		return err
	}

	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success get user data",
		Data:    data,
	})
}
