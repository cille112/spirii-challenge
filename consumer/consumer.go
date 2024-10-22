package consumer

import (
	"code-challenge/models"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func ConsumeStream(dataStream chan models.Data, db *sql.DB) {
	for data := range dataStream {
		if !checkForAnomalies(data) {
			err := storeInDB(db, data)
			if err != nil {
				log.Fatal("Could not store data - stopping")
			}
			fmt.Printf("Received data from MeterID: %s | Time: %s | Consumer: %s | Reading: %d\n",
				data.MeterID, data.Timestamp.Format(time.RFC3339), data.ConsumerID, data.MeterReading)
		}
	}
}

func checkForAnomalies(data models.Data) bool {
	return data.MeterReading < 0
}

func storeInDB(db *sql.DB, data models.Data) error {
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM data WHERE timestamp = ? AND meterId = ?)`
	err := db.QueryRow(checkQuery, data.Timestamp, data.MeterID).Scan(&exists)
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
		return fmt.Errorf("Failed to insert data: %w", err)
	}

	return nil
}
