package memory

import (
	"context"
	"time"

	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/logging"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var db *gorm.DB

var toMigrate = []interface{}{
	&User{},
	&ExternalID{},
	&Message{},
	&APIKey{},
}

func Setup(ctx context.Context) (*gorm.DB, error) {
	logger := logging.FromContext(ctx)

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
		return nil, errors.Wrap(errors.ErrOpenDatabase, err)
	}

	sqlDB, _ := db.DB()

	sqlDB.SetMaxIdleConns(config.DB.MaxIdle)
	sqlDB.SetMaxOpenConns(config.DB.MaxOpen)
	sqlDB.SetConnMaxLifetime(time.Hour)

	for k := range toMigrate {
		if err := db.AutoMigrate(toMigrate[k]); err != nil {
			return db, errors.Wrap(errors.ErrMigrateDatabase, err)
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
