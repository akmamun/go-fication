package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type DatabaseConfiguration struct {
	Driver   string
	Dbname   string
	Username string
	Password string
	Host     string
	Port     string
	LogMode  bool
}

func GetDNSConfig() string {
	masterName := viper.GetString("MASTER_DB_NAME")
	masterUser := viper.GetString("MASTER_DB_USER")
	masterPassword := viper.GetString("MASTER_DB_PASSWORD")
	masterHost := viper.GetString("MASTER_DB_HOST")
	masterPort := viper.GetString("MASTER_DB_PORT")
	masterSslMode := viper.GetString("MASTER_SSL_MODE")

	//replicaName := viper.GetString("REPLICA_DB_NAME")
	//replicaUser := viper.GetString("REPLICA_DB_USER")
	//replicaPassword := viper.GetString("REPLICA_DB_PASSWORD")
	//replicaHost := viper.GetString("REPLICA_DB_HOST")
	//replicaPort := viper.GetString("REPLICA_DB_PORT")
	//replicaSslMode := viper.GetString("REPLICA_SSL_MODE")

	masterDSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		masterHost, masterUser, masterPassword, masterName, masterPort, masterSslMode,
	)
	//
	//replicaDSN := fmt.Sprintf(
	//	"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
	//	replicaHost, replicaUser, replicaPassword, replicaName, replicaPort, replicaSslMode,
	//)
	return masterDSN
}
