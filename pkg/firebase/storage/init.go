package storage

import (
	"context"
	"log"
	"os"

	"firebase.google.com/go/v4/storage"
	"github.com/rhyth-me/backend/pkg/firebase"
)

var (
	// Client - firebase firestore client
	Client *storage.Client
	// Temporary bucket name
	Temp string
	// Image bucket name
	Image string
)

// Init - setup cloud firestore client
func Init() *storage.Client {
	ctx := context.Background()
	app := firebase.Init()

	cli, err := app.Storage(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	Client = cli

	Temp = os.Getenv("TEMPORARY_BUCKET")
	Image = os.Getenv("IMAGE_BUCKET")

	return cli
}
