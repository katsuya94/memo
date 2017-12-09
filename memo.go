package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/katsuya94/memo/auth"
	"google.golang.org/api/option"
)

const dateFmt = "2006-01-02"

func main() {
	ctx := context.Background()

	tokenSource, err := auth.TokenSourceFromConfig(ctx)
	if err != nil {
		fmt.Printf("Failed to get OAuth token: %v", err)
		os.Exit(1)
	}

	client, err := storage.NewClient(ctx, option.WithTokenSource(tokenSource))
	if err != nil {
		fmt.Printf("Failed to create client: %v", err)
		os.Exit(1)
	}

	bucket := client.Bucket("atateno-notes")
	object := bucket.Object(time.Now().Format(dateFmt))

	w := object.NewWriter(ctx)

	fmt.Fprintln(w, "Hello, World!")

	if err := w.Close(); err != nil {
		log.Fatalf("Failed to close bucket writer: %v", err)
	}
}
