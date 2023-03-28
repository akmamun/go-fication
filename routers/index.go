package routers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func RegisterRoutes(router *chi.Mux) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("\"live\": \"ok\""))
	})
	//Add All route
	ExamplesRoutes(router)
}
