package repository

import (
	"chi-boilerplate/helpers/pagination"
	"chi-boilerplate/models"
)

type ExampleRepo interface {
	GetExamples(limit, offset int64) (res interface{}, err error)
	CreateExample(exp *models.Example) error
}

func (r *GormRepository) GetExamples(limit, offset int64) (res interface{}, err error) {
	var example []*models.Example
	res = pagination.Paginate(&pagination.Param{
		DB:      r.db,
		Limit:   limit,
		Offset:  offset,
		OrderBy: "id ASC",
	}, &example)
	return
}
func (r *GormRepository) GetExamplesList() (exp []*models.Example, err error) {
	err = r.db.Database.Find(&exp).Error
	return
}

func (r *GormRepository) CreateExample(exp *models.Example) (err error) {
	err = r.db.Database.Create(exp).Error
	return
}
