package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	var err error
	DB, err = sql.Open("postgres", "user=youruser dbname=yourdb sslmode=disable")
	if err != nil {
		return err
	}
	return DB.Ping()
}
func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
func GetDB() *sql.DB {
	if DB == nil {
		panic("Database connection is not initialized")
	}
	return DB
}
