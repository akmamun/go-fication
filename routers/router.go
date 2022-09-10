package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoute() *chi.Mux {

	//allowedHosts := viper.GetString("ALLOWED_HOSTS")

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	//router.Use(cors.CORSMiddleware())
	//router.Use(cors.Handler(cors.Options{
	//	// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
	//	AllowedOrigins: []string{"*"},
	//	// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
	//	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	//	ExposedHeaders:   []string{"Link"},
	//	AllowCredentials: false,
	//	MaxAge:           300, // Maximum value not ignored by any of major browsers
	//}))
	//router.Route("/", func(r chi.Router) {
	//	r.Mount("/brands", (ctrl))
	//})
	//RegisterRoutes(handle, router)

	return router
}
