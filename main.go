package main

import (
	"code-challenge/consumer"
	"code-challenge/db"
	"code-challenge/models"
	"code-challenge/router"
	"code-challenge/stream"
	"log"

	"github.com/gin-gonic/gin"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "code-challenge/docs"
)

func main() {
	db := db.InitializeDatabase()
	defer db.Close()

	dataStream := make(chan models.Data)

	go stream.SimulateStream(dataStream)

	go consumer.ConsumeStream(dataStream, db)

	r := gin.Default()

	r.GET("/swagger/*any", gin.WrapH(httpSwagger.WrapHandler))

	r.GET("/data", router.BasicAuth(), func(c *gin.Context) {
		router.SensorDataHandler(c.Writer, c.Request, db)
	})

	r.GET("/topconsumer", router.BasicAuth(), func(c *gin.Context) {
		router.TopConsumerHandler(c.Writer, c.Request, db)
	})

	r.GET("/thirtyconsumer", router.BasicAuth(), func(c *gin.Context) {
		router.ThirthyPercentHandler(c.Writer, c.Request, db)
	})

	// r.GET("/forecast", router.BasicAuth(), func(c *gin.Context) {
	// 	router.ForecastHandler(c.Writer, c.Request, db)
	// })

	log.Fatal(r.Run(":8000"))

}
