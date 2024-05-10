package database_test

import (
	"testing"

	"github.com/golangcimri/worker/environ"
	"github.com/golangcimri/worker/globals"
	"github.com/golangcimri/worker/models"
)

// TestBulkInsertRecords tests the bulk insert functionality of the database.
func TestBulkInsertRecords(t *testing.T) {

	// Enviroment initialize
	environ.Init("test")

	// Records for testing
	records := []models.Record{
		{ID: 1, Title: "Record 1"},
		{ID: 2, Title: "Record 2"},
	}

	// Insert records
	db := globals.Variables.Database
	err := db.BulkInsertRecords(db.Collection.Tests, records)
	if err != nil {
		t.Fatalf("Failed to insert records: %v", err)
	}

	// Check the inserted records
	result := db.GetOne(db.Collection.Tests, 1)
	if result.Err() != nil {
		t.Fatalf("Failed to query records: %v", err)
	}

	// Decode the record

	var record models.Record
	err = result.Decode(&record)
	if err != nil {
		t.Fatalf("Failed to decode record: %v", err)
	}

	// Check the result
	if records[0].ID != record.ID || records[0].Title != record.Title {
		t.Fatalf("Expected record %+v, got %+v", records[0], record)
	}

}
