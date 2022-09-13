package routers

import (
	"github.com/go-chi/chi/v5"
	"go-fication/controllers"
	"go-fication/infra/database"
	"go-fication/repository"
)

func ExamplesRoutes(router *chi.Mux, db *database.DB) {
	repo := repository.NewGormRepository(db)
	exampleCtrl := controllers.NewExampleHandler(repo)
	router.Group(func(r chi.Router) {
		r.Get("/test", exampleCtrl.GetData)
		r.Post("/test", exampleCtrl.CreateData)

	})
}
