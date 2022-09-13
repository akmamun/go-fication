package main

import (
	"github.com/spf13/viper"
	"go-fication/config"
	"go-fication/infra/database"
	"go-fication/infra/logger"
	"go-fication/routers"
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

	db, err := database.DBConnection(config.GetDNSConfig())
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
