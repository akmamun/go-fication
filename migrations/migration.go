package migrations

import (
	"go-fication/models"
	"go-fication/repository"
)

func Migrate(db *repository.DB) {
	var migrationModels = []interface{}{&models.Example{}}
	err := db.Database.AutoMigrate(migrationModels...)
	if err != nil {
		return
	}
}
