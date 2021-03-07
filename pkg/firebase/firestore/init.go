package firestore

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/golang/glog"
	"github.com/rhyth-me/backend/pkg/firebase"
)

// Client - firebase firestore client
var Client *firestore.Client

var (
	// Users Collection name - Store user information.
	Users string
	// Items Collection name - Store item information.
	Items string
)

// Init - setup cloud firestore client
func Init() *firestore.Client {
	ctx := context.Background()
	app := firebase.Init()

	cli, err := app.Firestore(ctx)
	if err != nil {
		glog.Errorln(err)
	}

	Client = cli

	Users = os.Getenv("USERS_COLLECTION")
	Items = os.Getenv("ITEMS_COLLECTION")

	return cli
}
