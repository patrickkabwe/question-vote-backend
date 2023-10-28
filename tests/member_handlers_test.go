package tests

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"vote-app/internal/handlers"
	"vote-app/internal/models"
	"vote-app/internal/routes"
	"vote-app/internal/services"
	"vote-app/mocks"
	"vote-app/types"
	"vote-app/utils"
)

func TestMemberHandler_Register(t *testing.T) {
	testCases := []struct {
		name       string
		statusCode int
		body       []byte
		response   models.Member
	}{
		{
			name:       "valid request",
			statusCode: 200,
			body:       []byte(`{"email":"test@gmail.com","password":"123456"}`),
			response:   models.Member{Email: "test@gmail.com", Name: "test", ID: 1},
		},
		{
			name:       "invalid request",
			statusCode: 400,
			body:       []byte(`{"email":"test@gmail.com","password":"123"}`),
			response:   models.Member{Email: "", Name: "", ID: 0},
		},
	}

	app := utils.CreateApp()
	memberRepo := mocks.NewRepo()
	memberService := services.NewMemberService(memberRepo)
	memberHandler := handlers.NewMemberHandler(memberService)
	routes.MemberRoutes(memberHandler, app)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/members/login", bytes.NewReader(tc.body))
			req.Header.Set("Content-Type", "application/json")
			var bodyReq types.LoginRequest
			err := json.Unmarshal(tc.body, &bodyReq)
			if err != nil {
				t.Fatal(err)
			}
			filter := map[string]interface{}{"email": bodyReq.Email}
			memberRepo.On("FindOne", filter).Return(tc.response, nil)

			res, _ := app.Test(req)
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					t.Fatal(err)
				}
			}(res.Body)

			var response models.Member
			body, _ := io.ReadAll(res.Body)

			if err := json.Unmarshal(body, &response); err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, http.StatusOK, res.StatusCode)
			assert.Equal(t, tc.response.ID, response.ID)
		})
	}

}
