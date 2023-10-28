package routes

import (
	"github.com/gofiber/fiber/v2"
	"vote-app/database"
	"vote-app/internal/handlers"
	"vote-app/internal/services"
)

func VoteRoutes(db *database.DB, app *fiber.App) {
	// register resources
	voteService := services.NewVoteService(db)
	voteHandler := handlers.NewVoteHandler(voteService)

	//setup group
	votes := app.Group("/votes")

	// setup routes
	votes.Get("/", voteHandler.GetVote)
	votes.Post("/", voteHandler.CreateVote)
	votes.Delete("/", voteHandler.DeleteVote)

}
