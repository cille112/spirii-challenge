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
			storeInDB(db, data)
			fmt.Printf("Received data from MeterID: %s | Time: %s | Consumer: %s | Reading: %d\n",
				data.MeterID, data.Timestamp.Format(time.RFC3339), data.ConsumerID, data.MeterReading)
		}
	}
}

func checkForAnomalies(data models.Data) bool {
	return data.MeterReading < 0
}

func storeInDB(db *sql.DB, data models.Data) {
	query := `INSERT INTO data(meterId, timestamp, consumerId, meterReading) VALUES(?, ?, ?, ?)`
	_, err := db.Exec(query, data.MeterID, data.Timestamp.Format(time.RFC3339), data.ConsumerID, data.MeterReading)
	if err != nil {
		log.Fatalf("Failed to insert data: %v", err)
	}
}
