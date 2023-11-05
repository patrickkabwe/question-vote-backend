package handlers

import (
	"github.com/gofiber/fiber/v2"
	"vote-app/constants"
	"vote-app/internal/models"
	"vote-app/internal/services"
	"vote-app/types"
	"vote-app/utils"
)

type MemberHandler struct {
	memberService services.MemberService
}

func NewMemberHandler(memberService services.MemberService) MemberHandler {
	return MemberHandler{
		memberService: memberService,
	}
}

func (m *MemberHandler) Login(c *fiber.Ctx) error {
	var loginRequest types.LoginRequest
	member := models.Member{}
	if err := c.BodyParser(&loginRequest); err != nil {
		return c.JSON(types.APIResponse[*models.MemberResponse]{
			Success: false,
			Error:   fiber.ErrUnprocessableEntity.Error(),
			Data:    member.ToResponse(),
		})
	}

	if err := utils.LoginRequestValidate(&loginRequest); err != nil {
		return c.JSON(types.APIResponse[*models.MemberResponse]{
			Success: false,
			Error:   err.Error(),
			Data:    member.ToResponse(),
		})
	}

	foundUser, err := m.memberService.FindByEmail(loginRequest.Email)
	if err != nil {
		return c.JSON(types.APIResponse[*models.MemberResponse]{
			Success: false,
			Error:   err.Error(),
			Data:    member.ToResponse(),
		})
	}
	c.Cookie(&fiber.Cookie{
		Name:     constants.TOKEN_NAME,
		Value:    utils.GenerateToken(foundUser.ID),
		HTTPOnly: true,
	})
	return c.JSON(types.APIResponse[*models.MemberResponse]{
		Success: true,
		Error:   "",
		Data:    foundUser.ToResponse(),
	})
}

func (m *MemberHandler) Register(c *fiber.Ctx) error {
	var registerRequest types.RegisterRequest

	if err := c.BodyParser(&registerRequest); err != nil {
		return err
	}

	return c.JSON(registerRequest)
}
