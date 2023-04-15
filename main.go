package main

import (
	"be-ifid/adapter"
	"be-ifid/config"
	"be-ifid/database"
	"be-ifid/internal"
	"be-ifid/internal/handler"
	"be-ifid/internal/service"
	"be-ifid/utils"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func init() {
	config.Init()
	database.PostgresInit()
	database.RedisInit()
	adapter.MQTTInit()
	adapter.EmailInit()
}

func main() {
	conf := config.Get()
	ctx := context.Background()
	postgres := database.PostgresConnection()
	redis := database.RedisConnection(ctx)

	defer func() {
		err := postgres.Close()
		if err != nil {
			utils.Error("Error closing postgres database: ", err)
		}
		err = redis.Close()
		if err != nil {
			utils.Error("Error closing redis database: ", err)
		}
	}()

	app := fiber.New(config.FiberConfig())
	app.Use(recover.New())
	app.Use(cors.New(config.CorsConfig()))
	app.Use(logger.New())

	mqtt := adapter.MQTTConnection()
	service.NewMQTTService(*mqtt, conf).Subscibe()

	route := internal.NewRouter(app)

	baseHandler := handler.NewBaseHandler()
	authHandler := handler.NewAuthHandler(ctx, conf, postgres, redis)

	route.BaseRoute(baseHandler)
	route.AuthRoute(authHandler)

	baseUrl := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	err := app.Listen(baseUrl)
	if err != nil {
		utils.Error("Error app serve: ", err)
	}
}
