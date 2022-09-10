package repository

import (
	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) GormRepository {
	return GormRepository{
		db: db,
	}
}
