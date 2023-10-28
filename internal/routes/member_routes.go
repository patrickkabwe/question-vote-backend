package routes

import (
	"github.com/gofiber/fiber/v2"
	"vote-app/internal/handlers"
)

func MemberRoutes(handler handlers.MemberHandler, app *fiber.App) {

	//setup group
	members := app.Group("/members")

	// setup routes
	members.Post("/login", handler.Login)
	members.Post("/register", handler.Register)
}
