package tests_test

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"vote-app/constants"
	"vote-app/internal/handlers"
	"vote-app/internal/models"
	"vote-app/internal/routes"
	"vote-app/internal/services"
	"vote-app/mocks"
	"vote-app/types"
	"vote-app/utils"
)

func TestMemberHandler_Login(t *testing.T) {
	_ = os.Setenv("GO_ENV", "test")
	testCases := []struct {
		name       string
		statusCode int
		body       []byte
		response   types.APIResponse[models.Member]
	}{
		{
			name:       "POST /members/login success with valid email and password",
			statusCode: fiber.StatusOK,
			body:       []byte(`{"email":"test@gmail.com","password":"123456"}`),
			response: types.APIResponse[models.Member]{
				Success: true,
				Error:   "",
				Data:    models.Member{Email: "test@gmail.com", Name: "test", ID: 1},
			},
		},
		{
			name:       "POST /members/login failed with invalid email",
			statusCode: fiber.StatusBadRequest,
			body:       []byte(`{"email":"","password":"123"}`),
			response: types.APIResponse[models.Member]{
				Success: false,
				Error:   constants.ErrEmailRequired.Error(),
				Data:    models.Member{},
			},
		},
		{
			name:       "POST /members/login failed with empty password",
			statusCode: fiber.StatusBadRequest,
			body:       []byte(`{"email":"test@gmail.com","password":""}`),
			response: types.APIResponse[models.Member]{
				Success: false,
				Error:   constants.ErrPasswordRequired.Error(),
				Data:    models.Member{},
			},
		},
	}

	for _, tc := range testCases {
		app := utils.CreateApp()
		memberRepo := mocks.NewRepo()
		memberService := services.NewMemberService(memberRepo)
		memberHandler := handlers.NewMemberHandler(memberService)
		routes.MemberRoutes(memberHandler, app)

		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/members/login", bytes.NewReader(tc.body))
			req.Header.Set("Content-Type", "application/json")
			var bodyReq types.LoginRequest
			err := json.Unmarshal(tc.body, &bodyReq)
			if err != nil {
				t.Fatal(err)
			}
			filter := map[string]interface{}{"email": bodyReq.Email}

			if tc.response.Success {
				memberRepo.On("FindOne", filter).Return(tc.response.Data, nil)
			} else {
				memberRepo.On("FindOne", filter).Return(models.Member{}, tc.response.Error)
			}

			res, _ := app.Test(req)
			defer func(Body io.ReadCloser) {
				err = Body.Close()
				if err != nil {
					t.Fatal(err)
				}
			}(res.Body)

			var response types.APIResponse[models.Member]
			body, _ := io.ReadAll(res.Body)

			if err = json.Unmarshal(body, &response); err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, http.StatusOK, res.StatusCode)
			assert.Equal(t, tc.response, response)
			if tc.response.Success {
				assert.Contains(t, res.Header.Get("Set-Cookie"), constants.TOKEN_NAME)
			} else {
				assert.NotContains(t, res.Header.Get("Set-Cookie"), constants.TOKEN_NAME)
				assert.Equal(t, tc.response.Error, response.Error)
				assert.Equal(t, tc.response.Data.ToResponse(), response.Data.ToResponse())
			}
		})
	}
}
