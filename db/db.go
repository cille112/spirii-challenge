package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func InitializeDatabase() *sql.DB {
	// Open (or create) the SQLite database file
	db, err := sql.Open("sqlite", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create a table if it doesn't exist
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS data (
        meterId TEXT,
		consumerId TEXT,
        timestamp DATETIME,
		meterReading INT        
    );`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	return db
}
