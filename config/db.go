package config

type DBDriver string

const (
	Postgres  DBDriver = "postgres"
	MySQL     DBDriver = "mysql"
	SQLite    DBDriver = "sqlite"
	SQLServer DBDriver = "sqlserver"
)

type DBConfig struct {
	Driver  DBDriver
	DSN     string
	MaxIdle int
	MaxOpen int
}

var DB = &DBConfig{
	Driver:  DBDriver(envOrDefault("CHLOE_DB_DRIVER", string(SQLite))),
	DSN:     envOrDefault("CHLOE_DB_DSN", "chloe.db"), // if using SQLite, you can use 'file::memory:?cache=shared' for in-memory database
	MaxIdle: envOrDefaultInt("CHLOE_DB_MAX_IDLE", 2),
	MaxOpen: envOrDefaultInt("CHLOE_DB_MAX_OPEN", 10),
}
