package utils

import (
	"github.com/hibakun/nova-store/database"
	"github.com/hibakun/nova-store/models"
	"net/mail"
)

func GetUserByEmail(email string) (*models.User, bool) {
	var user models.User
	if err := database.DB.Where("email = ?", email).Find(&user); err.RowsAffected < 1 {
		return nil, false
	}
	return &user, true
}

func GetUserByPhoneNumber(number string) (*models.User, bool) {
	var user models.User
	if err := database.DB.Where("phone_number = ?", number).Find(&user); err.RowsAffected < 1 {
		return nil, false
	}
	return &user, true
}

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
