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