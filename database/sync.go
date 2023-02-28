package database

import (
	"github.com/hibakun/nova-store/models/model"
)

func Sync() {
	err := DB.AutoMigrate(
		&model.User{},
		&model.Game{},
		&model.Genre{},
		&model.Item{},
		&model.Payment{},
	)
	if err != nil {
		panic("failed to migrate database")
	}
}
