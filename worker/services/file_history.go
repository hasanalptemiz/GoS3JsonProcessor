package services

import (
	"fmt"
	"os"
)

// Insert a new object key to the history file when the object is downloaded successfully
func AppendToHistory(objectKey string) error {
	// Open the history file
	file, err := os.OpenFile("history.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Append the object key to the history file
	if _, err := fmt.Fprintf(file, "%s\n", objectKey); err != nil {
		return err
	}

	return nil
}
