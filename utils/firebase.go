package utils

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/golang/glog"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func initFirebase() *firebase.App {
	godotenv.Load(".env")

	opt := option.WithCredentialsJSON([]byte(os.Getenv("SERVICE_ACCOUNT_KEY")))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		glog.Errorln("Error")
	}
	return app
}

// InitFirestore - setup firestore client
func InitFirestore() *firestore.Client {
	app := initFirebase()
	firestore, err := app.Firestore(context.Background())
	if err != nil {
		glog.Errorln(err)
	}
	FS := firestore
	return FS
}
