package repository

import (
	"go-fication/infra/logger"
)

type GormRepository struct {
	db *DB
}

func NewGormRepository() *GormRepository {
	db, err := DBConnection()
	if err != nil {
		logger.Fatal("%v", err)
	}
	return &GormRepository{
		db: db,
	}
}
