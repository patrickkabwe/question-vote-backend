package services

import (
	"vote-app/internal/models"
	"vote-app/types"
)

type MemberService interface {
	FindByEmail(email string) (models.Member, error)
	FindById(id string) (models.Member, error)
}

type memberService struct {
	memberRepo types.Repository
}

func NewMemberService(repo types.Repository) MemberService {
	return &memberService{
		memberRepo: repo,
	}
}

func (m memberService) FindByEmail(email string) (models.Member, error) {
	filter := map[string]interface{}{"email": email}
	user, err := m.memberRepo.FindOne(filter)
	if err != nil {
		return models.Member{}, err
	}
	return user.(models.Member), nil
}

func (m memberService) FindById(id string) (models.Member, error) {
	//TODO implement me
	panic("implement me")
}
