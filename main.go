package main

import (
	"github.com/spf13/viper"
	"go-fication/config"
	"go-fication/infra/database"
	"go-fication/infra/logger"
	"go-fication/migrations"
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
	logLevel := viper.GetString("LOG_LEVEL")
	logger.SetLogLevel(logLevel)

	db, err := database.DBConnection(config.GetDNSConfig())
	if err != nil {
		logger.Fatal("%v", err)
	}
	//later remove auto migration
	migrations.Migrate(db)

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
