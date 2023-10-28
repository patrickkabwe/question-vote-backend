package services

import (
	"vote-app/types"
)

type QuestionService interface {
	CreateQuestion() error
	UpdateQuestion() error
	DeleteQuestion() error
	GetQuestion() error
	GetQuestions() error
}

type questionService struct {
	questionRepo types.Repository
}

func NewQuestionService(questionRepo types.Repository) QuestionService {
	return &questionService{
		questionRepo: questionRepo,
	}
}

func (q questionService) CreateQuestion() error {
	//TODO implement me
	panic("implement me")
}

func (q questionService) UpdateQuestion() error {
	//TODO implement me
	panic("implement me")
}

func (q questionService) DeleteQuestion() error {
	//TODO implement me
	panic("implement me")
}

func (q questionService) GetQuestion() error {
	//TODO implement me
	panic("implement me")
}

func (q questionService) GetQuestions() error {
	//TODO implement me
	panic("implement me")
}
