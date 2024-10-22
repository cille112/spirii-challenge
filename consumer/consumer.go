package consumer

import (
	database "code-challenge/db"
	"code-challenge/models"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func ConsumeStream(dataStream chan models.Data, db *sql.DB) {
	for data := range dataStream {
		if !checkForAnomalies(data) {
			err := database.StoreInDB(db, data)
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
