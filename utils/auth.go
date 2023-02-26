package utils

import (
	"github.com/hibakun/nova-store/database"
	"github.com/hibakun/nova-store/models/model"
	"net/mail"
)

func GetUserByEmail(email string) (*model.User, bool) {
	var user model.User
	if err := database.DB.Where("email = ?", email).Find(&user); err.RowsAffected < 1 {
		return nil, false
	}
	return &user, true
}

func GetUserByPhoneNumber(number string) (*model.User, bool) {
	var user model.User
	if err := database.DB.Where("phone_number = ?", number).Find(&user); err.RowsAffected < 1 {
		return nil, false
	}
	return &user, true
}

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
