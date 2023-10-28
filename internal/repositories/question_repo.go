package repositories

import (
	"vote-app/database"
	"vote-app/types"
)

type QuestionRepo struct {
	db *database.DB
}

func NewQuestionRepo(db *database.DB) types.Repo {
	return &QuestionRepo{
		db: db,
	}
}

func (q *QuestionRepo) Save(data any) (any, error) {
	//TODO implement me
	return QuestionRepo{}, nil
}

func (q *QuestionRepo) Update(id string, data any) (any, error) {
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
