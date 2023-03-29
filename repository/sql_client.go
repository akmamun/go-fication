package repository

import (
	"fmt"
	"github.com/spf13/viper"
	logger "go-fication/infra/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLog "gorm.io/gorm/logger"
	"time"
)

type DB struct {
	Database *gorm.DB
}

func DBConnection() (*DB, error) {
	masterDBDSN := fmt.Sprintf(
		"host=postgres_db user=dbuser password=password123 dbname=test_gofication_db port=5432 sslmode=disable")
	logMode := viper.GetBool("MASTER_DB_LOG_MODE")
	loglevel := gormLog.Silent
	if logMode {
		loglevel = gormLog.Info
	}
	pgConn := postgres.New(postgres.Config{
		DSN:                  masterDBDSN,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	})

	db, err := gorm.Open(pgConn, &gorm.Config{Logger: gormLog.Default.LogMode(loglevel)})

	if err != nil {
		logger.Fatal("database refused %v", err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetConnMaxIdleTime(time.Minute * 1)
	//sqlDB.SetMaxIdleConns(10)
	//sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Minute * 5)
	err = sqlDB.Ping()
	if err != nil {
		logger.Fatal("%v", err)
	}
	logger.Log("database connected")

	return &DB{Database: db}, nil
}

//if !debug {
//	//	database.Use(dbresolver.Register(dbresolver.Config{
//	//		Replicas: []gorm.Dialector{
//	//			postgres.Open(replicaDSN),
//	//		},
//	//		Policy: dbresolver.RandomPolicy{},
//	//	}))
