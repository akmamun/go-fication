package repository

import (
	"chi-boilerplate/models"
)

type ExampleRepo interface {
	GetExamples() (exp []*models.Example, err error)
	CreateExample(exp *models.Example) error
}

func (r GormRepository) GetExamples() (exps []*models.Example, err error) {
	err = r.db.Database.Find(&exps).Error
	return
}

func (r GormRepository) CreateExample(exp *models.Example) (err error) {
	err = r.db.Database.Create(exp).Error
	return
}
