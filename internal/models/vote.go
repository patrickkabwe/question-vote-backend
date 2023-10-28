package models

type Vote struct {
	ID         int64  `json:"id"`
	QuestionID string `json:"question_id"`
	MemberID   string `json:"member_id"`
}
