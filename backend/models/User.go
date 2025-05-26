
package models

import "time"

type User struct {
	ID                 string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    firstName               string    `gorm:"not null"`
    lastName               string    `gorm:"not null"`
    username string 
	Email              string    `gorm:"uniqueIndex;not null"`
	HashedPassword     string    `gorm:"not null"`
	GoogleAccessToken  string
	GoogleRefreshToken string
	TokenExpiry        time.Time
	CreatedAt          time.Time `gorm:"autoCreateTime"`
}
