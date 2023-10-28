package main

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"vote-app/docs"
)

func CreateApp() *fiber.App {
	app := fiber.New()

	app.Use(logger.New(
		logger.Config{
			Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
		}))

	app.Use(swagger.New(docs.Config))

	return app
}
