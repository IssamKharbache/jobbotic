package models

import (
	"time"
)

type JobApplication struct {
	ID        string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    string `gorm:"not null"`             // Foreign key to the user
	EmailID   string `gorm:"uniqueIndex;not null"` // Gmail message ID
	Subject   string
	From      string
	To        string
	Snippet   string
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
