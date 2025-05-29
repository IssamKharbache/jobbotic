
package utils

import (
	"jobbotic-backend/database"
	"jobbotic-backend/models"
	"time"
	"strconv"
)

func SaveGoogleToken(userID string, accessToken, refreshToken string, expiry time.Time) error {
	db := database.DB

	uid, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return err
	}

	var user models.User
	if err := db.First(&user, uid).Error; err != nil {
		return err
	}

	user.GoogleAccessToken = accessToken
	user.GoogleRefreshToken = refreshToken
	user.TokenExpiry = expiry
	user.IsGmailLinked = true

	return db.Save(&user).Error
}
