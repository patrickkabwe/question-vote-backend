package services

import "vote-app/database"

type VoteService interface {
	CreateVote() error
	DeleteVote() error
	GetVote() error
}

type voteService struct {
	db *database.Database
}

func NewVoteService(db *database.Database) VoteService {
	return &voteService{
		db: db,
	}
}

func (v voteService) CreateVote() error {
	//TODO implement me
	panic("implement me")
}

func (v voteService) DeleteVote() error {
	//TODO implement me
	panic("implement me")
}

func (v voteService) GetVote() error {
	//TODO implement me
	panic("implement me")
}
