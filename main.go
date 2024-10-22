package main

import (
	"code-challenge/consumer"
	"code-challenge/models"
	"code-challenge/stream"
)

func main() {

	dataStream := make(chan models.Data)

	go stream.SimulateStream(dataStream)

	consumer.ConsumeStream(dataStream)
}
