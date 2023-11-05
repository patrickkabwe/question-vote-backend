package services

import (
	"errors"
	"gorm.io/gorm"
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

func (m *memberService) FindByEmail(email string) (models.Member, error) {
	filter := map[string]interface{}{"email": email}
	user, err := m.memberRepo.FindOne(filter)
	return m.checkError(user, err)
}

func (m *memberService) FindById(id string) (models.Member, error) {
	user, err := m.memberRepo.FindById(id)
	return m.checkError(user, err)
}

func (m *memberService) checkError(user any, err error) (models.Member, error) {
	switch {
	case err == nil:
		return user.(models.Member), nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return models.Member{}, errors.New("user not found")
	default:
		return models.Member{}, errors.New("something went wrong")
	}
}
