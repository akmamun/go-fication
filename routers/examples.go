package routers

import (
	"github.com/go-chi/chi/v5"
	"go-fication/controllers"
)

func ExamplesRoutes(router *chi.Mux) {
	exampleCtrl := controllers.ExampleHandler{}
	router.Group(func(r chi.Router) {
		r.Get("/test/", exampleCtrl.GetData)
		r.Post("/test/", exampleCtrl.CreateData)

	})
}
