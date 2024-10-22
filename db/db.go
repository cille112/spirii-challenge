package db

import (
	"code-challenge/models"
	"database/sql"
	"fmt"
	"log"
	"time"

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

func StoreInDB(db *sql.DB, data models.Data) error {
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM data WHERE timestamp = ? AND meterId = ?)`
	err := db.QueryRow(checkQuery, data.Timestamp.Format(time.RFC3339), data.MeterID).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check for duplicates: %w", err)
	}

	if exists {
		log.Println("Duplicate entry found. Skipping insertion.")
		return nil
	}

	query := `INSERT INTO data(meterId, timestamp, consumerId, meterReading) VALUES(?, ?, ?, ?)`
	_, err = db.Exec(query, data.MeterID, data.Timestamp.Format(time.RFC3339), data.ConsumerID, data.MeterReading)
	if err != nil {
		return fmt.Errorf("failed to insert data: %w", err)
	}

	return nil
}

func GetData(db *sql.DB) ([]models.Data, error) {
	rows, err := db.Query("SELECT meterId, timestamp, consumerId, meterReading FROM data")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sensorData []models.Data
	for rows.Next() {
		var data models.Data
		if err := rows.Scan(&data.MeterID, &data.Timestamp, &data.ConsumerID, &data.MeterReading); err != nil {
			return nil, err
		}
		sensorData = append(sensorData, data)
	}
	return sensorData, nil
}

func GetTopConsumers(db *sql.DB, timeLimit string) ([]models.TopConsumer, error) {
	query := `
        SELECT consumerId, SUM(meterReading) AS totalReading
        FROM data
		WHERE timestamp >= ?
        GROUP BY consumerId
        ORDER BY totalReading DESC`

	// Prepare the statement
	rows, err := db.Query(query, timeLimit)
	if err != nil {
		return nil, fmt.Errorf("failed to query top consumers: %w", err)
	}
	defer rows.Close()

	// Slice to hold the results
	var topConsumers []models.TopConsumer

	// Iterate through the results
	for rows.Next() {
		var consumer models.TopConsumer
		if err := rows.Scan(&consumer.ConsumerID, &consumer.TotalReading); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		topConsumers = append(topConsumers, consumer)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while iterating rows: %w", err)
	}

	return topConsumers, nil
}

func GetTotalConsumption(db *sql.DB) (int, error) {

	query := `SELECT SUM(meterReading) FROM data`

	var totalConsumption sql.NullInt64
	err := db.QueryRow(query).Scan(&totalConsumption)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate total consumption: %w", err)
	}
	return int(totalConsumption.Int64), nil
}
