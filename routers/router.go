package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go-fication/infra/database"
	"go-fication/routers/middlewares"
)

func SetupRoute(db *database.DB) *chi.Mux {

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middlewares.Cors())

	RegisterRoutes(router, db)
	return router
}
