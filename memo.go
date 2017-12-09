package main

import (
	"context"
	"fmt"
	"io/ioutil"
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

	r, err := object.NewReader(ctx)
	if err == storage.ErrObjectNotExist {
		edit([]byte(""))
	} else if err == nil {
		initial, err := ioutil.ReadAll(r)
		if err != nil {
			fmt.Printf("Failed to read memo: %v", err)
		}

		edit(initial)
	} else {
		fmt.Printf("Failed to read memo: %v", err)
		os.Exit(1)
	}
}

func edit(initial []byte) []byte {
	return initial
}
