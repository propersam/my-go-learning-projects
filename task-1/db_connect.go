package main

import (
	"database/sql"
	"fmt"
	"log"
)

func createDbPool() (db *sql.DB) {
	// connecting to database
	dsn := fmt.Sprintf("%s:%s@/%s", dbUser, dbPass, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// db.SetMaxIdleConns()
	// db.SetMaxOpenConns()
	// db.SetConnMaxLifetime()

	return
}
