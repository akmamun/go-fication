package main

import (
	"github.com/spf13/viper"
	"go-fication/config"
	"go-fication/infra/logger"
	"go-fication/repository"
	"go-fication/routers"
	"net/http"
	"time"
)

func main() {
	viper.SetDefault("SERVER_TIMEZONE", "Asia/Dhaka")
	viper.SetDefault("LOG_LEVEL", "DEBUG")
	logLevel := viper.GetString("LOG_LEVEL")

	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc

	logger.SetLogLevel(logLevel)

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	//db, err := repository.DBConnection(config.GetDNSConfig())
	//if err != nil {
	//	logger.Fatal("%v", err)
	//}
	//later remove auto migration
	//migrations.Migrate(db)

	repo := repository.NewGormRepository() //future pass all config to repository
	routeHandler := routers.SetupRoute()
	serverConfig := &http.Server{
		Addr:              config.HttpServerConfig(),
		Handler:           routeHandler,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}
	httpServer := config.HttpServer{Server: serverConfig, Cfg: cfg, PgRepository: repo}
	logger.Fatal("%v", httpServer.ListenAndServe())
}
