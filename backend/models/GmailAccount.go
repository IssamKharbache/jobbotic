package models

import "time"

type GmailAccount struct {
	ID     string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID string `gorm:"type:uuid;not null;index"`    // foreign key to User
	User   User   `gorm:"constraint:OnDelete:CASCADE"` // optional, for eager loading

	Email              string    `gorm:"uniqueIndex;not null"` // Gmail address
	GoogleAccessToken  string    `gorm:"not null"`
	GoogleRefreshToken string    `gorm:"not null"`
	TokenExpiry        time.Time `gorm:"not null"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
