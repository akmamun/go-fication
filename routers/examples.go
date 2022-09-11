package routers

import (
	"chi-boilerplate/controllers"
	"chi-boilerplate/infra/database"
	"chi-boilerplate/repository"
	"github.com/go-chi/chi/v5"
)

func ExamplesRoutes(router *chi.Mux, db *database.DB) {
	repo := repository.NewGormRepository(db)
	exmCrtl := controllers.NewHandler(repo)
	router.Group(func(r chi.Router) {
		r.Get("/test", exmCrtl.GetData)
		r.Post("/test", exmCrtl.CreateData)

	})
}
