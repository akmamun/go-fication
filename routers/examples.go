package routers

import (
	"chi-boilerplate/controllers"
	"chi-boilerplate/infra/database"
	"chi-boilerplate/repository"
	"github.com/go-chi/chi/v5"
)

func ExamplesRoutes(router *chi.Mux, db *database.DB) {
	repo := repository.NewGormRepository(db)
	exampleCtrl := controllers.NewExampleHandler(repo)
	router.Group(func(r chi.Router) {
		r.Get("/test", exampleCtrl.GetData)
		r.Post("/test", exampleCtrl.CreateData)

	})
}
