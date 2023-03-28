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
	Driver: DBDriver(envOrDefaultWithOptions("CHLOE_DB_DRIVER", string(SQLite),
		[]string{string(Postgres), string(MySQL), string(SQLite), string(SQLServer)})),
	DSN:     envOrDefault("CHLOE_DB_DSN", "chloe.db"), // if using SQLite, you can use 'file::memory:?cache=shared' for in-memory database
	MaxIdle: envOrDefaultIntInRange("CHLOE_DB_MAX_IDLE", 2, 0, 100),
	MaxOpen: envOrDefaultIntInRange("CHLOE_DB_MAX_OPEN", 10, 1, 1000),
}
