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

	router := routers.SetupRoute(db)
	server := http.Server{
		Addr:    config.ServerConfig(),
		Handler: router,
		//ErrorLog: logger.DefaultErrLogger,
		//ReadTimeout:  cfg.ReadTimeout,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	logger.Fatalf("%v", server.ListenAndServe())
}
