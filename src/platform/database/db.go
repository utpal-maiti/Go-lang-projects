package database

import (
	"database/sql"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// DB is the global database connection pool.
var DB *sql.DB

// Connect establishes a connection to the PostgreSQL database and pings it.
func Connect() error {
	var err error
	DB, err = sql.Open("postgres", "user=youruser dbname=yourdb sslmode=disable")
	if err != nil {
		return err
	}
	return DB.Ping()
}