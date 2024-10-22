package stream

import (
	"code-challenge/models"
	"math/rand/v2"
	"time"
)

var meterIdList = []string{"sensor-1", "sensor-2", "sensor-3", "sensor-4", "sensor-5", ""}
var consumerIdList = []string{"consumer-1", "consumer-2", "consumer-3", "consumer-4", "consumer-5"}

func SimulateStream(dataStream chan models.Data) {
	for {
		data := generateRandomData(meterIdList, consumerIdList)
		dataStream <- data
		randomInterval := time.Duration(500+rand.IntN(2500)) * time.Millisecond
		time.Sleep(randomInterval)

		//Sometimes data is sent twice
		if rand.Float64() < 0.1 {
			dataStream <- data
			randomInterval := time.Duration(500+rand.IntN(2500)) * time.Millisecond
			time.Sleep(randomInterval)
		}
	}
}

func generateRandomData(meterIdList []string, consumerIdList []string) models.Data {
	randomMeterID := meterIdList[rand.IntN(len(meterIdList))]
	randomConsumerID := consumerIdList[rand.IntN(len(consumerIdList))]
	return models.Data{
		Timestamp:    time.Now(),
		MeterID:      randomMeterID,
		ConsumerID:   randomConsumerID,
		MeterReading: generateRandomReading(200),
	}
}

func generateRandomReading(max int) int {
	num := rand.IntN(max)

	//Sometimes it's negative
	if rand.Float64() < 0.1 {
		num = -num
	}

	return num
}
