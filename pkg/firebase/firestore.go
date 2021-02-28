package firebase

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/golang/glog"
)

// InitFirestore - setup firestore client
func InitFirestore() *firestore.Client {
	app := initFirebase()
	firestore, err := app.Firestore(context.Background())
	if err != nil {
		glog.Errorln(err)
	}
	return firestore
}
