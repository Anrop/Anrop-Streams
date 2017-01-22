package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Database *sql.DB
)

func Connect(databaseURL string) error {
	var err error
	Database, err = sql.Open("mysql", databaseURL)
	return err
}
