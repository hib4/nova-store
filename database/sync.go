package database

import (
	"github.com/hibakun/nova-store/models"
)

func Sync() {
	err := DB.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		panic("failed to migrate database")
	}
}
