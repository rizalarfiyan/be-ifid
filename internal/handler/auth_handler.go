package handler

import "github.com/gofiber/fiber/v2"

type AuthHandler interface {
	Login(ctx *fiber.Ctx) error
	Callback(ctx *fiber.Ctx) error
	FirstUser(ctx *fiber.Ctx) error
	Me(ctx *fiber.Ctx) error
}
