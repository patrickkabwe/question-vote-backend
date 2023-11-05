package utils

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"vote-app/config"
	"vote-app/docs"
)

func CreateApp() *fiber.App {
	app := fiber.New()

	app.Use(logger.New(
		logger.Config{
			Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
		}))
	if config.GetEnv("GO_ENV") != "test" {
		app.Use(swagger.New(docs.Config))
	}

	return app
}
