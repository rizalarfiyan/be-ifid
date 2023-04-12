package main

import (
	"be-ifid/config"
	"be-ifid/database"
	"be-ifid/internal"
	"be-ifid/internal/handler"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func init() {
	config.Init()
	database.Init()
}

func main() {
	conf := config.Get()

	app := fiber.New(config.FiberConfig())
	app.Use(recover.New())
	app.Use(cors.New(config.CorsConfig()))
	app.Use(logger.New())

	baseHandler := handler.NewBaseHandler()

	route := internal.NewRouter(app)
	route.BaseRoute(baseHandler)

	baseUrl := fmt.Sprintf("%s:%d", conf.ServerHost, conf.ServerPort)
	err := app.Listen(baseUrl)
	if err != nil {
		log.Fatalf("Error app serve: %v \n", err.Error())
	}
}
