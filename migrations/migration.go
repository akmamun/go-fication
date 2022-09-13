package migrations

import (
	"go-fication/infra/database"
	"go-fication/models"
)

func Migrate(db *database.DB) {
	var migrationModels = []interface{}{&models.Example{}}
	err := db.Database.AutoMigrate(migrationModels...)
	if err != nil {
		return
	}
}
