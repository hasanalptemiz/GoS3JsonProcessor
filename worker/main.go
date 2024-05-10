package main

import (
	"log"
	"sync"

	"github.com/golangcimri/worker/environ"
	"github.com/golangcimri/worker/globals"
	"github.com/golangcimri/worker/services"
	"github.com/golangcimri/worker/storage"
)

func Init() {
	environ.Init("main")
}

func main() {
	Init()
	objectKeys := []string{"products-1.jsonl", "products-2.jsonl", "products-3.jsonl", "products-4.jsonl"}
	// Create a group of workers
	var wg sync.WaitGroup

	db := globals.Variables.Database
	// For each object key, download the file from S3, process the data, and insert it into MongoDB
	for _, objectKey := range objectKeys {

		// Check if the objectKey is processed before
		if services.IsProcessed(objectKey) {
			log.Printf("%s file has already been processed", objectKey)
			continue
		}

		wg.Add(1)
		go func(objectKey string) {
			defer wg.Done()

			// Get the file data from S3
			s3Service, err := storage.NewS3(globals.Variables.AccessKey, globals.Variables.SecretKey, globals.Variables.Region)
			if err != nil {
				log.Printf("error creating s3 session: %v", err)
				return
			}

			// Download the object
			fileData, err := s3Service.DownloadObject(globals.Variables.BucketName, objectKey)
			if err != nil {
				log.Printf("error dowloanding object: %v", err)
				return
			}

			// Process the data and get the records
			records, err := services.ProcessData(fileData)
			if err != nil {
				log.Printf("error processing the data and get the records for %s: %v", objectKey, err)
				return
			}

			// Insert the records into MongoDB
			err = db.BulkInsertRecords(db.Collection.Records, records)
			if err != nil {
				log.Printf("error inserting MongoDB: %v", err)
				return
			}

			// Append the object key to the history
			err = services.AppendToHistory(objectKey)
			if err != nil {
				log.Printf("error appending the object key to the history: %v", err)
				return
			}

			log.Printf("%s file is successfuly writed", objectKey)
		}(objectKey)
	}

	// Tüm işlemlerin tamamlanmasını bekle
	wg.Wait()
}
