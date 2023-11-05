package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"vote-app/internal/models"
	"vote-app/internal/repositories"
)

func TestMemberRepo_Create(t *testing.T) {
	testCases := []struct {
		name    string
		payload models.Member
		willErr bool
	}{
		{
			name: "Should create member success",
			payload: models.Member{
				Name:     "test",
				Email:    "test@gmail.com",
				Password: "123456",
			},
			willErr: false,
		},
		{
			name: "Should fail to create member with empty name",
			payload: models.Member{
				Name:     "",
				Email:    "test@gmail.com",
				Password: "123456",
			},
			willErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db := setupDB()
			memberRepo := repositories.NewMemberRepo(db)
			user, err := memberRepo.Create(tc.payload)
			if tc.willErr {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, tc.payload.Name, user.(models.Member).Name)
				assert.Nil(t, err)
			}
		})
	}
}

func TestMemberRepo_Update(t *testing.T) {
	testCases := []struct {
		name    string
		arg     int64
		payload models.Member
		willErr bool
	}{
		{
			name: "Should update member success",
			arg:  1,
			payload: models.Member{
				Name: "test",
			},
			willErr: false,
		},
		{
			name: "Should fail to create member with empty name",
			arg:  100,
			payload: models.Member{
				Name:     "",
				Email:    "test@gmail.com",
				Password: "123456",
			},
			willErr: true,
		},
	}

	oldUser := models.Member{
		Name:     "test1",
		Email:    "test2@gmail.com",
		Password: "123456",
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db := setupDB()
			memberRepo := repositories.NewMemberRepo(db)
			if tc.arg == 1 {
				_, err := memberRepo.Create(oldUser)
				if err != nil {
					t.Fatal(err)
				}
			}
			user, err := memberRepo.Update(tc.arg, tc.payload)

			if !tc.willErr {
				assert.Nil(t, err)
			}
			assert.NotEqual(t, oldUser.Name, user.(models.Member).Name)
		})
	}
}
