# Database Configuration

Chloe uses GORM as its ORM, so you can use any database supported by GORM.

Currently, the drivers for PostgreSQL, MySQL, SQLite, and SQL Server are implemented, but you
can [ask for another driver](https://github.com/kamushadenes/chloe/issues/new?assignees=kamushadenes&labels=feature&template=feature_request.md&title=%5BFEATURE%5D+)
as the implementation should be trivial.

| Environment Variable | Default Value | Description                                                                                           | Options                                     |
|----------------------|---------------|-------------------------------------------------------------------------------------------------------|---------------------------------------------|
| CHLOE_DB_DRIVER      | sqlite        | Database driver to use                                                                                | postgres<br/>mysql<br/>sqlite<br/>sqlserver |
| CHLOE_DB_DSN         | chloe.db      | Database connection string, refer to the [docs](https://gorm.io/docs/connecting_to_the_database.html) |                                             |
| CHLOE_DB_MAX_IDLE    | 2             | Maximum number of idle connections                                                                    |                                             |
| CHLOE_DB_MAX_OPEN    | 10            | Maximum number of open connections                                                                    |                                             |
