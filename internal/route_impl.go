package internal

import (
	"be-ifid/internal/handler"
	"be-ifid/middleware"

	"github.com/gofiber/fiber/v2"
)

type router struct {
	app *fiber.App
}

func NewRouter(app *fiber.App) Router {
	return &router{
		app: app,
	}
}

func (r *router) BaseRoute(handler handler.BaseHandler) {
	r.app.Get("/", handler.Home)
}

func (r *router) AuthRoute(handler handler.AuthHandler) {
	auth := r.app.Group("/auth")
	auth.Post("/login", handler.Login)
	auth.Get("/callback", handler.Callback)

	protected := auth.Group("/", middleware.NewJWTMiddleware(middleware.JWTConfig{}))
	protected.Get("/me", handler.Me)
	protected.Post("/first-user", handler.FirstUser)
}
