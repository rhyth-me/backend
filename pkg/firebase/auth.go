package firebase

import (
	"context"

	"firebase.google.com/go/auth"
	"github.com/golang/glog"
)

// InitAuth - setup firebase auth client
func InitAuth() *auth.Client {
	app := initFirebase()
	auth, err := app.Auth(context.Background())
	if err != nil {
		glog.Errorln(err)
	}
	return auth
}
