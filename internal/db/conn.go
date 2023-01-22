package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Connect connects to the database
func Connect() error {
	var err error
	db, err = sql.Open("mysql", "u351287655_root:pKYS;0830@tcp(sql861.main-hosting.eu)/u351287655_kys")
	if err != nil {
		return err
	}
	fmt.Println("Success login into database!")
	return db.Ping()
}

// Close closes the database connection
func Close() {
	db.Close()
}
