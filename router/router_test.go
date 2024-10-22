package router

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestBasicAuth(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Set up a route with BasicAuth middleware
	router.GET("/protected", BasicAuth(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	// Create a test request without Authorization header
	req := httptest.NewRequest("GET", "/protected", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status %d, got %d", http.StatusUnauthorized, w.Code)
	}

	// Create a test request with invalid credentials
	invalidCredentials := "invalid:credentials"
	encodedCredentials := base64.StdEncoding.EncodeToString([]byte(invalidCredentials))
	req.Header.Set("Authorization", "Basic "+encodedCredentials)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status %d, got %d", http.StatusUnauthorized, w.Code)
	}

	//Try with valid credentials
	validCredentials := username + ":" + password
	encodedCredentials = base64.StdEncoding.EncodeToString([]byte(validCredentials))
	req.Header.Set("Authorization", "Basic "+encodedCredentials)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}
