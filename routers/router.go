package routers

import (
	"chi-boilerplate/infra/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoute(db *database.DB) *chi.Mux {

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(Cors())

	RegisterRoutes(router, db)
	return router
}
