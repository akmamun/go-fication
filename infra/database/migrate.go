package database

import (
	"embed"
	"errors"
	"fmt"
	"go-fication/infra/logger"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/viper"

	// migrate tools
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

const (
	_defaultAttempts = 20
	_defaultTimeout  = time.Second
)

var (
	attempts = _defaultAttempts
	err      error
	m        *migrate.Migrate
)

//go:embed migrations/*.sql
var fs embed.FS

func Migrate() {
	databaseURL := viper.GetString("MASTER_DB_URL")
	// if !ok || len(databaseURL) == 0 {
	// 	logger.Fatal("migrate: environment variable not declared: PG_URL")
	// }
	fmt.Println(databaseURL)
	d, err := iofs.New(fs, "migrations")
	if err != nil {
		logger.Fatal(err.Error())
	}
	m, err := migrate.NewWithSourceInstance("iofs", d, databaseURL)
	if err != nil {
		logger.Fatal(err.Error())
	}

	// for attempts > 0 {
	// 	m, err = migrate.New("file://infra/database/migrations", databaseURL)
	// 	if err == nil {
	// 		break
	// 	}

	// 	logger.Log("Migrate: postgres is trying to connect, attempts left: %d", attempts)
	// 	time.Sleep(_defaultTimeout)
	// 	attempts--
	// }

	if err != nil {
		logger.Fatal("Migrate: postgres connect error: %s", err)
	}

	err = m.Up()
	defer m.Close()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Fatal("Migrate: up error: %s", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		logger.Log("Migrate: no change")
		return
	}

	logger.Log("Migrate: up success")
}

// func Migrate(db *database.DB) {
// 	var migrationModels = []interface{}{
// 		&models.Transaction{},
// 		&models.TransactionLog{},
// 		&models.KgdclPayment{},
// 	}
// 	err := db.Database.AutoMigrate(migrationModels...)
// 	if err != nil {
// 		return
// 	}
// }
