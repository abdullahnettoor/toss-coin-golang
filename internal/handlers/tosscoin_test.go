// handlers_test.go

package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestTossCoin(t *testing.T) {
	// Create a new Gin router
	router := gin.New()

	// Define a route for the TossCoin function
	router.GET("/toss-coin", TossCoin)

	// Create a mock HTTP request to the /toss-coin endpoint
	req, err := http.NewRequest("GET", "/toss-coin", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the request to the router
	router.ServeHTTP(recorder, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, recorder.Code, "Status code should be 200")

	// Parse the JSON response
	var response map[string]bool
	assert.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &response), "Failed to parse JSON response")

	// Check the response structure
	assert.Contains(t, response, "res", "Response should contain 'res' field")
	assert.IsType(t, false, response["res"], "Response 'res' field should be a boolean")
	
	// You may add more specific assertions based on the expected behavior of your TossCoin function.
	// For example, you can check if the "res" field is a boolean or if it represents the expected result.
}
