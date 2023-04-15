package internal

import "be-ifid/internal/handler"

type Router interface {
	BaseRoute(handler handler.BaseHandler)
	AuthRoute(handler handler.AuthHandler)
}
