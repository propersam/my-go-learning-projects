package main

import (
	"database/sql"
	"fmt"
	"log"
)

func openDbPool() sql.DB {
	// connecting to database
	dsn := fmt.Sprintf("%s:%s@/%s", dbUser, dbPass, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return *db
	// db.SetMaxIdleConns()
	// db.SetMaxOpenConns()
	// db.SetConnMaxLifetime()
}
