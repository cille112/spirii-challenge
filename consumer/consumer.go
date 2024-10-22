package consumer

import (
	"code-challenge/models"
	"fmt"
	"time"
)

func ConsumeStream(dataStream chan models.Data) {
	for data := range dataStream {
		fmt.Printf("Received data from MeterID: %s | Time: %s | Consumer: %s | Reading: %d\n",
			data.MeterID, data.Timestamp.Format(time.RFC3339), data.ConsumerID, data.MeterReading)
	}
}
