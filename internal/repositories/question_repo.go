package repositories

import (
	"vote-app/database"
	"vote-app/types"
)

type QuestionRepo struct {
	db *database.Database
}

func NewQuestionRepo(db *database.Database) types.Repository {
	return &QuestionRepo{
		db: db,
	}
}

func (q *QuestionRepo) Create(data any) (any, error) {
	//TODO implement me
	return QuestionRepo{}, nil
}

func (q *QuestionRepo) Update(id int64, data any) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (q *QuestionRepo) Delete(id string) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (q *QuestionRepo) FindById(id string) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (q *QuestionRepo) FindOne(filter map[string]any) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (q *QuestionRepo) FindAll(filter map[string]any) (any, error) {
	//TODO implement me
	panic("implement me")
}
