package utils

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/golang/glog"
	"google.golang.org/api/option"
)

func initFirebase() *firebase.App {
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
	return firestore
}

// InitAuth - setup firebase auth client
func InitAuth() *auth.Client {
	app := initFirebase()
	auth, err := app.Auth(context.Background())
	if err != nil {
		glog.Errorln(err)
	}
	return auth
}
