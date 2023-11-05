package models

import "time"

type Member struct {
	ID       int64      `json:"id,omitempty" gorm:"primary_key;auto_increment:true;"`
	Name     string     `json:"name,omitempty" gorm:"type:varchar(255);not null;"`
	Email    string     `json:"email,omitempty" gorm:"type:varchar(255);not null;unique"`
	Password string     `json:"password,omitempty" gorm:"type:varchar(255);not null"`
	JoinedAt *time.Time `json:"joined_at,omitempty" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

type MemberResponse struct {
	ID       int64      `json:"id,omitempty"`
	Name     string     `json:"name,omitempty"`
	Email    string     `json:"email,omitempty"`
	JoinedAt *time.Time `json:"joined_at,omitempty"`
}

func (m *Member) TableName() string {
	return "members"
}

func (m *Member) ToResponse() *MemberResponse {
	if m == nil {
		return nil
	}
	if m.ID == 0 {
		return nil
	}
	return &MemberResponse{
		ID:       m.ID,
		Name:     m.Name,
		Email:    m.Email,
		JoinedAt: m.JoinedAt,
	}
}
