package config

import "time"

type Database struct {
	// Host is the database host.
	Host string `default:"127.0.0.1:5432" envconfig:"DB_HOST"`

	// Schema is the database schema.
	Schema string `default:"public" envconfig:"DB_SCHEMA"`

	// Database is the database name.
	Database string `default:"sumup" envconfig:"DB_DATABASE"`

	// Username is the db username.
	Username string `default:"sumup" envconfig:"DB_USERNAME"`

	// Password is the user password.
	Password string `default:"sumup" envconfig:"DB_PASSWORD"`

	// SSLMode is the SSL mode to use when connecting to Database.
	SSLMode string `default:"disable" envconfig:"DB_SSL_MODE"`

	// PingTimeout is the timeout in seconds for performing initial ping for the database.
	PingTimeout time.Duration `default:"5s" envconfig:"DB_PING_TIMEOUT"`

	// ConnectTimeout is the timeout for creating new connections.
	ConnectTimeout time.Duration `default:"5s" envconfig:"DB_CONNECT_TIMEOUT"`

	// TimeZone is the timezone to use for the SQL queries while connected to Database.
	Timezone string `default:"UTC" envconfig:"DB_TIMEZONE"`

	// MaxOpenConnections is the maximum number of open database connections.
	MaxOpenConnections int `default:"18" envconfig:"DB_MAX_OPEN_CONNECTIONS"`

	// MaxIdleConnections is the maximum number of idle database connections.
	MaxIdleConnections int `default:"18" envconfig:"DB_MAX_IDLE_CONNECTIONS"`

	// MaxConnLifetime is the maximum allowed connection lifetime.
	MaxConnLifetime time.Duration `default:"1h" envconfig:"DB_MAX_CONN_LIFETIME"`

	// MaxConnIdleTime is the maximum allowed time for a connection to stay in idle state.
	MaxConnIdleTime time.Duration `default:"30m" envconfig:"DB_MAX_CONN_IDLE_TIME"`

	// LogQueries specify whether to log SQL queries.
	//
	// NOTE: when LogQueries is false, no queries will be logged including the failures.
	LogQueries bool `default:"false" envconfig:"DB_LOG_QUERIES"`

	// LogSQL specify whether the query logs will log the SQL query.
	//
	// For example, when LogSQL is false, the query log won't have the "sql" field.
	LogSQL bool `default:"false" envconfig:"DB_LOG_SQL"`

	// LogArgs specify whether the query arguments must be logged.
	//
	// WARN: this MUST NOT be enabled on production.
	LogSQLArgs bool `default:"false" envconfig:"DB_LOG_SQL_ARGS"`
}
