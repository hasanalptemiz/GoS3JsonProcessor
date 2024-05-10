package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golangcimri/api/environ"
	"github.com/golangcimri/api/globals"
	"github.com/golangcimri/api/routes"
	"github.com/stretchr/testify/assert"
)

// Warning: You should write the records to the database before running the test.
// Integration test for the GetRecord function
func TestGetRecord(t *testing.T) {
	// Environment initialization
	environ.Init("test")

	// Fiber application setup
	app := fiber.New()

	// Route setup
	routes.PrivateRoutes(app)

	// Test request creation
	req := httptest.NewRequest(http.MethodGet, "/api/v1/private/get/record/1", nil)
	// Api-Token header is required
	req.Header.Set("Api-Token", globals.Variables.ApiToken)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	// Response status check
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Response body reading and checking
	var record map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&record)
	assert.NoError(t, err)
	assert.NotNil(t, record)

	// Check if the 'brand' field exists and has the expected value
	brandValue, brandExists := record["brand"].(string)
	assert.True(t, brandExists, "brand field should exist in the response")
	assert.Equal(t, "kron", brandValue, "brand field should have the expected value 'kron'")
}
