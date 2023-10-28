package models

import (
	"vote-app/database"
)

type Member struct {
	ID       int64  `json:"id" gorm:"primary_key;auto_increment:true;"`
	Name     string `json:"name" gorm:"type:varchar(255);not null;"`
	Email    string `json:"email" gorm:"type:varchar(255);not null;unique_index"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
	JoinedAt string `json:"joined_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	*database.DB
}
