package database

import (
	"fmt"
	"github.com/hibakun/nova-store/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Config("PGHOST"),
		config.Config("PGUSER"),
		config.Config("PGPASSWORD"),
		config.Config("PGDATABASE"),
		config.Config("PGPORT"),
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	Sync()
}
