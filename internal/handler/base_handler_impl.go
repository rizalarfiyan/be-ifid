package handler

import "github.com/gofiber/fiber/v2"

type BaseHandler interface {
	Home(ctx *fiber.Ctx) error
}
