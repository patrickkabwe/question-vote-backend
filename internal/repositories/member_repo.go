package repositories

import (
	"vote-app/database"
	"vote-app/internal/models"
	"vote-app/types"
)

type MemberRepo struct {
	db *database.DB
}

func NewMemberRepo(db *database.DB) types.Repository {
	return &MemberRepo{
		db: db,
	}
}

func (r *MemberRepo) Save(data any) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (r *MemberRepo) Update(id string, data any) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (r *MemberRepo) Delete(id string) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (r *MemberRepo) FindById(id string) (any, error) {
	var member models.Member
	results := r.db.Find(&member, id)
	return member, results.Error
}

func (r *MemberRepo) FindOne(filter map[string]any) (any, error) {
	var member models.Member
	results := r.db.Model(&member).Find(&member, filter)
	return member, results.Error
}

func (r *MemberRepo) FindAll(filter map[string]any) (any, error) {
	//TODO implement me
	panic("implement me")
}
