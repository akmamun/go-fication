package main

import (
	"chi-boilerplate/config"
	"chi-boilerplate/controllers"
	"chi-boilerplate/infra/database"
	"chi-boilerplate/infra/logger"
	"chi-boilerplate/repository"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func main() {
	l := logger.New(viper.GetString("Log_Level"))
	//set timezone
	viper.SetDefault("SERVER_TIMEZONE", "Asia/Dhaka")
	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc

	if err := config.SetupConfig(); err != nil {
		l.Fatal("config SetupConfig() error: %s", err)
	}

	db, err := database.DBConnection("postgres://mamun:123@localhost:5432/test_pg_go")
	if err != nil {
		logger.Fatalf("%v", err)
	}

	//TODO changes in later
	//migrations.Migrate(db.Database)
	//repo := repository.NewGormRepository(db)
	//handle := controllers.NewBaseHandler(repo)
	//router := routers.SetupRoute(r)

	//srvr := http.Server{
	//	Addr:    config.ServerConfig(),
	//	Handler: router,
	//	//ErrorLog: logger.DefaultErrLogger,
	//	//WriteTimeout: cfg.WriteTimeout,
	//	//ReadTimeout:  cfg.ReadTimeout,
	//	ReadTimeout:       5 * time.Second,
	//	WriteTimeout:      5 * time.Second,
	//	IdleTimeout:       30 * time.Second,
	//	ReadHeaderTimeout: 2 * time.Second,
	//}
	//if err := srvr.ListenAndServe(); err != http.ErrServerClosed {
	//	logger.Fatalf("bal ama")
	//}
	router := chi.NewRouter()
	repo := repository.NewGormRepository(db)
	exmCrtl := controllers.NewHandler(repo)
	router.Group(func(r chi.Router) {
		r.Get("/test", exmCrtl.GetData)
		r.Post("/test", exmCrtl.CreateData)

	})

	logger.Fatalf("%v", http.ListenAndServe(config.ServerConfig(), router))

}
