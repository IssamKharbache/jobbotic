package models

import "time"

type User struct {
	ID                 string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	FirstName          string    `json:"first_name" gorm:"not null"`
	LastName           string    `json:"last_name" gorm:"not null"`
	Email              string    `json:"email" gorm:"uniqueIndex;not null"`
	HashedPassword     string    `json:"-" gorm:"not null"` // omit from JSON responses
	GoogleAccessToken  string    `json:"-"`
	GoogleRefreshToken string    `json:"-"`
	TokenExpiry        time.Time `json:"-"`
	IsGmailLinked      bool      `json:"-"`
	LinkedEmail        *string   `json:"linked_email" gorm:"uniqueIndex:idx_users_linked_email"`
	CreatedAt          time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt          time.Time `json:"updatedAt" gorm:"autoCreateTime"`
}
