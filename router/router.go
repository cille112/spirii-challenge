package router

import (
	"code-challenge/models"
	"database/sql"
	"encoding/json"
	"net/http"
)

func SensorDataHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")

	data, err := getSensorData(db)
	if err != nil {
		http.Error(w, "Unable to fetch sensor data", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(data)
}

func getSensorData(db *sql.DB) ([]models.Data, error) {
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
