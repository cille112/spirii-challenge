package db

import (
	"code-challenge/models"
	"database/sql"
	"log"
	"testing"
	"time"
)

func setupTestDB() *sql.DB {
	// Create an in-memory database for testing
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	// Create the necessary table
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS data (
		meterId TEXT,
		consumerId TEXT,
		timestamp DATETIME,
		meterReading INT        
	);`

	if _, err := db.Exec(createTableQuery); err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	return db
}

func TestInitializeDatabase(t *testing.T) {
	// This function doesn't need to be tested extensively since it only sets up the database
	db := InitializeDatabase()
	defer db.Close()

	if db == nil {
		t.Error("Expected a database connection, got nil")
	}

	// Check if the table exists
	var count int
	row := db.QueryRow("SELECT count(*) FROM sqlite_master WHERE type='table' AND name='data'")
	err := row.Scan(&count)
	if err != nil {
		t.Errorf("Failed to check if table exists: %v", err)
	}
	if count != 1 {
		t.Errorf("Expected table 'data' to exist, but it does not")
	}
}

func TestStoreInDB(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	// Test data for insertion
	data := models.Data{
		Timestamp:    time.Now(),
		MeterID:      "meter1",
		ConsumerID:   "consumer1",
		MeterReading: 100,
	}

	// Attempt to store data
	err := StoreInDB(db, data)
	if err != nil {
		t.Fatalf("Failed to insert data: %v", err)
	}

	// Verify that the data has been inserted
	var meterID string
	row := db.QueryRow("SELECT meterId FROM data WHERE meterId = ?", data.MeterID)
	err = row.Scan(&meterID)
	if err != nil {
		t.Fatalf("Failed to retrieve inserted data: %v", err)
	}
	if meterID != data.MeterID {
		t.Errorf("Expected MeterID %s, got %s", data.MeterID, meterID)
	}

	// Attempt to insert the same data again (should be skipped)
	err = StoreInDB(db, data)
	if err != nil {
		t.Fatalf("Failed to insert data: %v", err)
	}

	// Check that only one entry exists in the database
	var count int
	row = db.QueryRow("SELECT COUNT(*) FROM data WHERE meterId = ?", data.MeterID)
	err = row.Scan(&count)
	if err != nil {
		t.Fatalf("Failed to count entries: %v", err)
	}
	if count != 1 {
		t.Errorf("Expected 1 entry for MeterID %s, found %d", data.MeterID, count)
	}
}

func TestGetData(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	data1 := models.Data{
		MeterID:      "meter-2",
		Timestamp:    time.Now().Add(time.Hour), // Adjust for time added
		ConsumerID:   "consumer2",
		MeterReading: 200,
	}

	data2 := models.Data{
		MeterID:      "meter-1",
		Timestamp:    time.Now(),
		ConsumerID:   "consumer1",
		MeterReading: 100,
	}

	err1 := StoreInDB(db, data1)

	err2 := StoreInDB(db, data2)

	if err1 != nil || err2 != nil {
		t.Errorf("expected no error, got %v", err1.Error())
		t.Errorf("expected no error, got %v", err2.Error())
	}

	data, err := GetData(db)

	// Check for errors
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Define expected data
	expectedData := []models.Data{
		data1,
		data2,
	}

	// Check the result
	if len(data) != len(expectedData) {
		t.Errorf("expected %d rows, got %d", len(expectedData), len(data))
		return
	}

	// Assert that the data returned is correct
	for i, v := range data {
		if v.MeterID != expectedData[i].MeterID || v.ConsumerID != expectedData[i].ConsumerID || v.MeterReading != expectedData[i].MeterReading {
			t.Errorf("expected %v, got %v", expectedData[i], v)
		}
	}
}
