package router

import (
	database "code-challenge/db"
	"code-challenge/models"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	username = "admin"
	password = "password"
)

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort() // Abort the request
			return
		}

		// Check if the authorization header is Basic
		if !strings.HasPrefix(auth, "Basic ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization type"})
			c.Abort()
			return
		}

		// Decode the base64 encoded credentials
		payload := strings.TrimPrefix(auth, "Basic ")
		decoded, err := base64.StdEncoding.DecodeString(payload)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid base64 encoding"})
			c.Abort()
			return
		}

		// Split the decoded credentials
		parts := strings.SplitN(string(decoded), ":", 2)
		if len(parts) != 2 || parts[0] != username || parts[1] != password {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// If authentication is successful, proceed to the next handler
		c.Next()
	}

}

// @Summary Get data
// @Description Get a list of data
// @ID get-data
// @Produce json
// @Success 200 {array} models.Data
// @Router /data [get]
func SensorDataHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")

	data, err := database.GetData(db)
	if err != nil {
		http.Error(w, "Unable to fetch sensor data", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(data)
}

// @Summary Get top consumers
// @Description Get a list of consumers and their usage for the last 10 minutes
// @ID get-top-consumer
// @Produce json
// @Success 200 {array} models.TopConsumer
// @Router /topconsumer [get]
func TopConsumerHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")

	data, err := database.GetTopConsumers(db, time.Now().Add(-10*time.Minute).Format(time.RFC3339))
	if err != nil {
		print(err.Error())
		http.Error(w, "Unable to fetch consumer data", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(data)
}

// @Summary Get consumers that adds up to around 30% to total consumption
// @Description Get a list of consumers and their usage that is almost 30% of total
// @ID get-tirthy-consumer
// @Produce json
// @Success 200 {array} models.TopThirtyConsumer
// @Router /thirtyconsumer [get]
func ThirthyPercentHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")

	data, err := getTop30Consumer(db)
	if err != nil {
		http.Error(w, "Unable to fetch consumer data", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(data)
}

func getTop30Consumer(db *sql.DB) (models.TopThirtyConsumer, error) {
	var resp models.TopThirtyConsumer

	totalConsumption, err := database.GetTotalConsumption(db)
	if err != nil {
		return resp, fmt.Errorf("failed to get total consumption: %w", err)
	}

	thirtyPercent := float64(totalConsumption) * 0.3

	tc, err := database.GetTopConsumers(db, time.Time{}.Format(time.RFC3339))
	if err != nil {
		return resp, fmt.Errorf("failed to get consumption pr consumer: %w", err)
	}

	// Sort consumers by TotalReading in descending order
	sort.Slice(tc, func(i, j int) bool {
		return tc[i].TotalReading > tc[j].TotalReading
	})

	// Accumulate consumer consumption until we reach or exceed the threshold
	var selectedConsumers []models.TopConsumer
	accumulated := 0
	for _, consumer := range tc {
		selectedConsumers = append(selectedConsumers, consumer)
		accumulated += consumer.TotalReading
		if accumulated >= int(thirtyPercent) {
			break
		}
	}

	resp.Consumers = selectedConsumers
	resp.TotalConsumption = totalConsumption

	return resp, nil
}
