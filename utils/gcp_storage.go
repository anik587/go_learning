package utils

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
)

func UploadGCP() {
	ctx := context.Background()

	// Set the environment variable for authentication
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./key.json")

	// Replace these variables with your bucket name and object name
	bucketName := "toffee_voucher"
	objectName := "test52.csv"

	// Create a pipe
	pr, pw := io.Pipe()
	writer := csv.NewWriter(pw)

	// Channel to capture any errors from the goroutine
	errChan := make(chan error, 1)

	go func() {
		defer pw.Close()
		defer writer.Flush()

		// Example CSV data
		data := [][]string{
			{"Header1", "Header2", "Header3"},
			{"Row1Col1", "Row1Col2", "Row1Col3"},
		}

		if err := writer.WriteAll(data); err != nil {
			errChan <- fmt.Errorf("failed to write data to CSV: %v", err)
			return
		}
		errChan <- nil
	}()

	// Initialize a GCS client
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create GCS client: %v", err)
	}
	defer client.Close()

	// Get a reference to the bucket and object
	bucket := client.Bucket(bucketName)
	obj := bucket.Object(objectName)

	// Create a writer to the GCS object
	wc := obj.NewWriter(ctx)
	defer wc.Close()

	// Copy the contents of the CSV data from the pipe to the GCS object
	if _, err := io.Copy(wc, pr); err != nil {
		log.Fatalf("Failed to upload file to GCS: %v", err)
	}

	// Check for any errors in the goroutine
	if err := <-errChan; err != nil {
		log.Fatalf("Error writing CSV data: %v", err)
	}

	// Generate the public URL of the uploaded file
	publicURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, objectName)

}
