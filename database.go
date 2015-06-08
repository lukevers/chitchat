package main

import (
	"database/sql"
	"fmt"
	"github.com/rubenv/sql-migrate"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strings"
)

var db *sql.DB

func dbconnect() {
	// Check to see if *database contains ?parseTime=true because if it does
	// not then we want to add it
	if !strings.Contains(*database, "?parseTime=true") {
		*database += "?parseTime=true"
	}

	// Connect to database
	db, err := sql.Open("mysql", *database)
	if err != nil {
		fmt.Printf("Error connecting to database: %s\n", err)
		os.Exit(1)
	}

	// Ping the database
	err = db.Ping()
	if err != nil {
		fmt.Printf("Error pinging database: %s\n", err)
		os.Exit(2)
	}

	// Setup migrations
	migrations := &migrate.FileMigrationSource{
		Dir: "database/migrations",
	}

	// Run migrations
	n, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
	if err != nil {
		fmt.Printf("Error running database migrations: %s\n", err)
		os.Exit(3)
	} else {
		if n == 0 {
			fmt.Println("Nothing to migrate")
		} else {
			fmt.Printf("Applied %d migrations\n", n)
		}
	}
}
