package main

import (
	"code-challenge/consumer"
	"code-challenge/db"
	"code-challenge/models"
	"code-challenge/router"
	"code-challenge/stream"
	"log"
	"net/http"
)

func main() {
	db := db.InitializeDatabase()
	defer db.Close()

	dataStream := make(chan models.Data)

	go stream.SimulateStream(dataStream)

	go consumer.ConsumeStream(dataStream, db)

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		router.SensorDataHandler(w, r, db)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))

}
