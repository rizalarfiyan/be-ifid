package main

import (
	"be-ifid/config"
	"be-ifid/database"
	"be-ifid/model"
	"fmt"
	"log"
	"net/http"

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

	app.Get("/", func(ctx *fiber.Ctx) error {
		ctx.JSON(model.BaseResponse{
			Code:    http.StatusOK,
			Message: "Success!",
			Data:    nil,
		})
		return nil
	})

	baseUrl := fmt.Sprintf("%s:%d", conf.ServerHost, conf.ServerPort)
	err := app.Listen(baseUrl)
	if err != nil {
		log.Fatalf("Error app serve: %v \n", err.Error())
	}
}
