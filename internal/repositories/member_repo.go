package repositories

import (
	"vote-app/database"
	"vote-app/internal/models"
	"vote-app/types"
)

type MemberRepo struct {
	db *database.Database
}

func NewMemberRepo(db *database.Database) types.Repository {
	return &MemberRepo{
		db: db,
	}
}

func (r *MemberRepo) Create(data any) (any, error) {
	member := models.Member{
		Name:     data.(models.Member).Name,
		Email:    data.(models.Member).Email,
		Password: data.(models.Member).Password,
	}

	dbTx := r.db.Create(&member)
	if dbTx.Error != nil {
		return nil, dbTx.Error
	}

	return member, nil
}

func (r *MemberRepo) Update(id int64, data any) (any, error) {

	member := models.Member{
		ID:    id,
		Name:  data.(models.Member).Name,
		Email: data.(models.Member).Email,
	}

	dbTx := r.db.Model(&member).Omit("id").Updates(member)
	if dbTx.Error != nil {
		return nil, dbTx.Error
	}

	return member, nil
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
	results := r.db.Select("id", "name", "email", "password").Find(&member, filter)
	return member, results.Error
}

func (r *MemberRepo) FindAll(filter map[string]any) (any, error) {
	//TODO implement me
	panic("implement me")
}
