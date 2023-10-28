package routes

import (
	"github.com/gofiber/fiber/v2"
	"vote-app/database"
	"vote-app/internal/handlers"
	"vote-app/internal/repositories"
	"vote-app/internal/services"
)

func QuestionRoutes(db *database.DB, app *fiber.App) {
	// register resources
	questionRepo := repositories.NewQuestionRepo(db)
	questionService := services.NewQuestionService(questionRepo)
	questionHandler := handlers.NewQuestionHandler(questionService)

	//setup group
	questions := app.Group("/questions")

	// setup routes
	questions.Get("/", questionHandler.GetQuestions)
	questions.Post("/", questionHandler.CreateQuestion)
	questions.Post("/:id", questionHandler.UpdateQuestion)
	questions.Delete("/:id", questionHandler.DeleteQuestion)
	questions.Get("/:id", questionHandler.GetQuestion)

}
