package handlers

import (
	"github.com/gofiber/fiber/v2"
	"vote-app/internal/services"
)

type QuestionHandler struct {
	questionService services.QuestionService
}

func NewQuestionHandler(questionService services.QuestionService) *QuestionHandler {
	return &QuestionHandler{
		questionService: questionService,
	}
}

func (q QuestionHandler) CreateQuestion(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (q QuestionHandler) UpdateQuestion(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (q QuestionHandler) DeleteQuestion(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (q QuestionHandler) GetQuestion(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (q QuestionHandler) GetQuestions(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
