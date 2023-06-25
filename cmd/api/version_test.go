package main_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/danstis/CrowdSound/internal/version"
	"github.com/stretchr/testify/assert"
)

// Test_GetVersion tests the "/version" endpoint of the API.
func Test_GetVersion(t *testing.T) {
	expected := version.Version

	// Create a new GET request to the "/version" endpoint.
	req, _ := http.NewRequest("GET", "/version", nil)

	// Execute the request and get the response.
	response := executeRequest(req)
	assert.Equal(t, http.StatusOK, response.Code)

	// Parse the response body into a map.
	var m map[string]interface{}
	_ = json.Unmarshal(response.Body.Bytes(), &m)

	assert.Equal(t, expected, m["version"])
}
