package database

func Sync() {
	err := DB.AutoMigrate()
	if err != nil {
		panic("failed to migrate database")
	}
}
