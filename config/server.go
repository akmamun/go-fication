package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go-fication/infra/logger"
	"go-fication/repository"
	"net/http"
)

type ServerConfiguration struct {
	Port                 string
	Secret               string
	LimitCountPerRequest int64
}

func HttpServerConfig() string {
	viper.SetDefault("SERVER_HOST", "0.0.0.0")
	viper.SetDefault("SERVER_PORT", "8000")

	httpServer := fmt.Sprintf("%s:%s", viper.GetString("SERVER_HOST"), viper.GetString("SERVER_PORT"))
	logger.Log("Server Running at %s:", httpServer)
	return httpServer
}

type HttpServer struct {
	*http.Server
	Cfg          *viper.Viper
	PgRepository *repository.GormRepository
}
