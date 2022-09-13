package main

import (
	"chi-boilerplate/config"
	"chi-boilerplate/infra/database"
	"chi-boilerplate/infra/logger"
	"chi-boilerplate/routers"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func main() {
	viper.SetDefault("SERVER_TIMEZONE", "Asia/Dhaka")
	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc

	if err := config.SetupConfig(); err != nil {
		logger.Error("%v", err)
	}

	db, err := database.DBConnection("postgres://mamun:123@localhost:5432/test_pg_go")
	if err != nil {
		logger.Fatal("%v", err)
	}

	router := routers.SetupRoute(db)
	server := http.Server{
		Addr:              config.ServerConfig(),
		Handler:           router,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	logger.Fatal("%v", server.ListenAndServe())
}
