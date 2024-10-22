package models

import "time"

// Data represents the structure of the data being fetched
// @Description Data structure for API
type Data struct {
	Timestamp    time.Time `json:"timestamp"`    // timestamp of data
	MeterID      string    `json:"meterId"`      // id of meter
	ConsumerID   string    `json:"consumerId"`   // id of consumer
	MeterReading int       `json:"meterReading"` // the reading from the meter
}

// TopConsumer struct to hold aggregated results
// @Description TopConsumer structure for API
type TopConsumer struct {
	ConsumerID   string `json:"consumerId"`
	TotalReading int    `json:"totalReading"`
}

// TopThirtyConsumer struct to hold aggregated results
// @Description TopThirtyConsumer structure for API
type TopThirtyConsumer struct {
	Consumers        []TopConsumer
	TotalConsumption int
}
