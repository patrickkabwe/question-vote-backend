package handlers

import (
	"github.com/gofiber/fiber/v2"
	"vote-app/internal/services"
	"vote-app/types"
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
	if err := c.BodyParser(&loginRequest); err != nil {
		return err
	}

	foundUser, err := m.memberService.FindByEmail(loginRequest.Email)
	if err != nil {
		return err
	}

	return c.JSON(foundUser)
}

func (m *MemberHandler) Register(c *fiber.Ctx) error {
	var registerRequest types.RegisterRequest

	if err := c.BodyParser(&registerRequest); err != nil {
		return err
	}

	return c.JSON(registerRequest)
}
