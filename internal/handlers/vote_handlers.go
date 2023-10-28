package handlers

import (
	"github.com/gofiber/fiber/v2"
	"vote-app/internal/services"
)

type VoteHandler interface {
	CreateVote(c *fiber.Ctx) error
	DeleteVote(c *fiber.Ctx) error
	GetVote(c *fiber.Ctx) error
}

type voteHandler struct {
	voteService services.VoteService
}

func NewVoteHandler(voteService services.VoteService) VoteHandler {
	return &voteHandler{
		voteService: voteService,
	}
}

func (v voteHandler) CreateVote(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (v voteHandler) DeleteVote(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (v voteHandler) GetVote(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
