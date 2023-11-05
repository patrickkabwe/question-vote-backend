package routes

import (
	"github.com/gofiber/fiber/v2"
	"vote-app/database"
	"vote-app/internal/handlers"
	"vote-app/internal/repositories"
	"vote-app/internal/services"
)

func SetupRoutes(db *database.Database, app *fiber.App) {
	// register resources
	memberRepo := repositories.NewMemberRepo(db)
	memberService := services.NewMemberService(memberRepo)
	memberHandler := handlers.NewMemberHandler(memberService)

	MemberRoutes(memberHandler, app)
	QuestionRoutes(db, app)
	VoteRoutes(db, app)
}
