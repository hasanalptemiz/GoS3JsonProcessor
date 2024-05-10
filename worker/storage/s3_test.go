package storage_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/golangcimri/worker/environ"
	"github.com/golangcimri/worker/globals"
	"github.com/golangcimri/worker/storage"
)

// TestDownloadObject tests the download functionality of the S3 service.
func TestDownloadObject(t *testing.T) {

	environ.Init("test")
	// Set up a temporary directory to store fake object data
	tmpDir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatalf("failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a new S3 service
	s3Service, err := storage.NewS3(globals.Variables.AccessKey, globals.Variables.SecretKey, globals.Variables.Region)
	if err != nil {
		t.Fatalf("failed to create S3 service: %v", err)
	}

	// Download the object
	downloadedData, err := s3Service.DownloadObject(globals.Variables.BucketName, "products-2.jsonl")
	if err != nil {
		t.Fatalf("failed to download object: %v", err)
	}

	// Test succeeded
	t.Logf("Downloaded data: %s", downloadedData)

}
