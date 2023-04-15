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
	adapter.MQTTInit()
	database.RedisInit()
}

func main() {
	conf := config.Get()
	ctx := context.Background()
	db := database.PostgresConnection()
	redis := database.RedisConnection(ctx)

	defer func() {
		err := db.Close()
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

	baseHandler := handler.NewBaseHandler()
	route := internal.NewRouter(app)
	route.BaseRoute(baseHandler)

	baseUrl := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	err := app.Listen(baseUrl)
	if err != nil {
		utils.Error("Error app serve: ", err)
	}
}
