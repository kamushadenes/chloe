package memory

import (
	"context"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/logging"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func Setup(ctx context.Context) (*gorm.DB, error) {
	logger := logging.GetLogger()

	logger.Info().Msg("initializing database")

	var err error

	switch config.DB.Driver {
	case config.Postgres:
		db, err = gorm.Open(postgres.Open(config.DB.DSN), &gorm.Config{
			Logger: DBLogger{},
		})
	case config.MySQL:
		db, err = gorm.Open(mysql.Open(config.DB.DSN), &gorm.Config{
			Logger: DBLogger{},
		})
	case config.SQLServer:
		db, err = gorm.Open(sqlserver.Open(config.DB.DSN), &gorm.Config{
			Logger: DBLogger{},
		})
	case config.SQLite:
		db, err = gorm.Open(sqlite.Open(config.DB.DSN), &gorm.Config{
			Logger: DBLogger{},
		})
	}

	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()

	sqlDB.SetMaxIdleConns(config.DB.MaxIdle)
	sqlDB.SetMaxOpenConns(config.DB.MaxOpen)
	sqlDB.SetConnMaxLifetime(time.Hour)

	var toMigrate []interface{}

	toMigrate = append(toMigrate, &User{})
	toMigrate = append(toMigrate, &ExternalID{})
	toMigrate = append(toMigrate, &Message{})
	toMigrate = append(toMigrate, &APIKey{})

	for _, model := range toMigrate {
		if err := db.AutoMigrate(model); err != nil {
			return db, err
		}
	}

	return db, nil
}

func Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			sqlDB, _ := db.DB()
			_ = sqlDB.Close()
		}
	}
}
