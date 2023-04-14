package main

import (
	"be-ifid/adapter"
	"be-ifid/config"
	"be-ifid/internal"
	"be-ifid/internal/handler"
	"be-ifid/internal/service"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func init() {
	config.Init()
	adapter.MQTTInit()
}

func main() {
	conf := config.Get()

	app := fiber.New(config.FiberConfig())
	app.Use(recover.New())
	app.Use(cors.New(config.CorsConfig()))
	app.Use(logger.New())

	mqtt := adapter.MQTTGet()
	service.NewMQTTService(*mqtt, conf).Subscibe()

	baseHandler := handler.NewBaseHandler()
	route := internal.NewRouter(app)
	route.BaseRoute(baseHandler)

	baseUrl := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	err := app.Listen(baseUrl)
	if err != nil {
		log.Fatalf("Error app serve: %v \n", err.Error())
	}
}
