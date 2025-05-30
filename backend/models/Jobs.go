package models

import (
	"time"
)

type JobApplication struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`             // Foreign key to the user
	EmailID   string `gorm:"uniqueIndex;not null"` // Gmail message ID
	Subject   string
	From      string
	To        string
	Snippet   string
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
