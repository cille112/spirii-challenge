package models

import "time"

type Data struct {
	Timestamp    time.Time
	MeterID      string
	ConsumerID   string
	MeterReading int
}
