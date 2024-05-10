package services

import (
	"bufio"
	"bytes"
	"encoding/json"
	"os"
	"strings"

	"github.com/golangcimri/worker/models"
)

// ProcessData processes the JSONL data and returns a slice of records
func ProcessData(data []byte) ([]models.Record, error) {
	var records []models.Record

	// Parse each line as a JSON object
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := scanner.Bytes()

		// Skip empty lines
		if len(line) == 0 {
			continue
		}

		// Parse JSON object into Record struct
		var record models.Record
		if err := json.Unmarshal(line, &record); err != nil {
			return nil, err
		}

		// Append record to records slice
		records = append(records, record)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

// isProcessed checks if the object key is already processed
func IsProcessed(objectKey string) bool {
	// Open the history file
	file, err := os.Open("history.txt")
	if err != nil {
		return false
	}
	defer file.Close()

	// Scan through the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == objectKey {
			// The objectKey is found in the history file
			return true
		}
	}

	// The objectKey is not found in the history file
	return false
}
