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
	// Should we use go routines to work concurrent if a lot of data?
	// Find a way to handle duplicate when they can be process by different to routines
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
	return data.MeterReading < 0 && data.MeterID == ""
}
