package memory

import (
	"context"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func Setup(ctx context.Context) (*gorm.DB, error) {
	logger := zerolog.Ctx(ctx)

	logger.Info().Msg("initializing database")

	var err error
	db, err = gorm.Open(sqlite.Open("chloe.db"), &gorm.Config{
		Logger: DBLogger{},
	})
	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	var toMigrate []interface{}

	toMigrate = append(toMigrate, &User{})
	toMigrate = append(toMigrate, &ExternalID{})
	toMigrate = append(toMigrate, &Message{})

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
