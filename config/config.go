package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Configuration struct {
	Server ServerConfiguration
}

func LoadConfig() (*viper.Viper, error) {

	config := viper.New()
	// Set the configuration file name and path
	config.SetConfigFile(".env")
	// Read the configuration file
	if err := config.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %s", err)
	}

	return config, nil
}

//SetupConfig configuration
//func SetupConfig() *Configuration {
//	var configuration *Configuration
//
//	viper.SetConfigFile(".env")
//	if err := viper.ReadInConfig(); err != nil {
//		logger.Error("Error to reading config file, %s", err)
//	}
//
//	err := viper.Unmarshal(&configuration)
//	if err != nil {
//		logger.Error("error to decode, %v", err)
//	}
//
//	return configuration
//}
