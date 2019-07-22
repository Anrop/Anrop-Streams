package db

import (
	"database/sql"

	// Set SQL Driver
	_ "github.com/go-sql-driver/mysql"
)

// Database is the connection holder for database access
var (
	Database *sql.DB
)

// Connect establishes a connection to the database server
func Connect(databaseURL string) error {
	var err error
	Database, err = sql.Open("mysql", databaseURL)
	return err
}
