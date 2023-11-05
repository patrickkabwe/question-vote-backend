package tests_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"vote-app/internal/models"
	"vote-app/internal/services"
	"vote-app/mocks"
)

func TestMemberService_FindByEmail(t *testing.T) {
	type args struct {
		email string
	}

	testCases := []struct {
		name     string
		args     args
		expected models.Member
		error    error
	}{
		{
			name: "should return member_test",
			args: args{
				email: "test@gmail.com",
			},
			expected: models.Member{
				ID:       1,
				Name:     "test",
				Email:    "test@gmail.com",
				Password: "hashed_password",
			},
			error: nil,
		},
		{
			name: "should return empty member_test",
			args: args{
				email: "",
			},
			expected: models.Member{},
			error:    gorm.ErrRecordNotFound,
		},
		{
			name:     "should return error",
			expected: models.Member{},
			error:    errors.New("something went wrong"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			memberRepo := mocks.NewRepo()
			memberService := services.NewMemberService(memberRepo)

			if tc.error != nil {
				memberRepo.On("FindOne", map[string]interface{}{"email": tc.args.email}).Return(tc.expected, tc.error)
				_, err := memberService.FindByEmail(tc.args.email)
				assert.Error(t, err)

				return
			}
			memberRepo.On("FindOne", map[string]interface{}{"email": tc.args.email}).Return(tc.expected, nil)

			member, err := memberService.FindByEmail(tc.args.email)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tc.expected, member)
		})
	}

}

func TestMemberService_FindById(t *testing.T) {
	type args struct {
		id string
	}

	testCases := []struct {
		name     string
		args     args
		expected models.Member
		error    error
	}{
		{
			name: "should return member_test",
			args: args{
				id: "1",
			},
			expected: models.Member{
				ID:       1,
				Name:     "test",
				Email:    "test@gmail.com",
				Password: "hashed_password",
			},
			error: nil,
		},
		{
			name: "should return empty member_test",
			args: args{
				id: "0",
			},
			expected: models.Member{},
			error:    gorm.ErrRecordNotFound,
		},
		{
			name:     "should return error",
			expected: models.Member{},
			error:    errors.New("something went wrong"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			memberRepo := mocks.NewRepo()
			memberService := services.NewMemberService(memberRepo)

			if tc.error != nil {
				memberRepo.On("FindById", tc.args.id).Return(tc.expected, tc.error)
				_, err := memberService.FindById(tc.args.id)
				assert.Error(t, err)

				return
			}
			memberRepo.On("FindById", tc.args.id).Return(tc.expected, nil)

			member, err := memberService.FindById(tc.args.id)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tc.expected, member)
		})
	}

}
