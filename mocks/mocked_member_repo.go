package mocks

import (
	"github.com/stretchr/testify/mock"
	"vote-app/internal/models"
)

// MockedMemberRepo is an autogenerated mock type for the Repo type
type MockedMemberRepo struct {
	mock.Mock
}

func NewRepo() *MockedMemberRepo {
	return &MockedMemberRepo{}
}

func (r *MockedMemberRepo) Create(data any) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (r *MockedMemberRepo) Update(id int64, data any) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (r *MockedMemberRepo) Delete(id string) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (r *MockedMemberRepo) FindById(id string) (any, error) {
	args := r.Called(id)
	return args.Get(0).(models.Member), args.Error(1)
}

func (r *MockedMemberRepo) FindOne(filter map[string]any) (any, error) {
	args := r.Called(filter)
	return args.Get(0).(models.Member), args.Error(1)
}

func (r *MockedMemberRepo) FindAll(filter map[string]any) (any, error) {
	//TODO implement me
	panic("implement me")
}
