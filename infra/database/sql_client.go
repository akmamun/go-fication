package database

import (
	logger "go-fication/infra/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type DB struct {
	Database *gorm.DB
}

func DBConnection(dsn string) (*DB, error) {

	//	debug := viper.GetBool("DEBUG")
	//	loglevel := logger.Silent
	//
	//
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		//Logger: log.Default.LogMode(logLevel),
	})

	if err != nil {
		logger.Fatal("%v", err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetConnMaxIdleTime(time.Minute * 1)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	err = sqlDB.Ping()
	if err != nil {
		logger.Fatal("%v", err)
	}

	return &DB{Database: db}, nil
}

//func ConnectDB() error {
//	//logMode := viper.GetBool("DB_LOG_MODE")
//	//debug := viper.GetBool("DEBUG")
//	//loglevel := logger.Silent
//	//
//	//if logMode {
//	//	loglevel = logger.Info
//	//}
//	database, err := gorm.Open(postgres.New(postgres.Config{
//		DSN:                  "postgres://mamun:123@localhost:5432/test_pg",
//		PreferSimpleProtocol: true, // disables implicit prepared statement usage
//	}), &gorm.Config{})
//
//	if err != nil {
//		return err
//	}
//	sqlDB, err := database.DB()
//	sqlDB.SetConnMaxIdleTime(time.Minute * 5)
//	sqlDB.SetMaxIdleConns(10)
//	sqlDB.SetMaxOpenConns(100)
//	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
//	sqlDB.SetConnMaxLifetime(time.Hour)
//	//postgres.Config{
//	//	DSN: masterDSN,
//	//}), &gorm.Config{
//	//	Logger: logger.Default.LogMode(loglevel),
//	//})
//	//p := models.Example{Data: "data"}
//	//database.Create(&p)
//
//	//if !debug {
//	//	database.Use(dbresolver.Register(dbresolver.Config{
//	//		Replicas: []gorm.Dialector{
//	//			postgres.Open(replicaDSN),
//	//		},
//	//		Policy: dbresolver.RandomPolicy{},
//	//	}))
//	//}
//	//defer sqlDB.Close()
//
//	//db := &DB{
//	//	Database: database,
//	//}
//	return nil
//}
