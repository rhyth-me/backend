package auth

import (
	"context"
	"os"

	"firebase.google.com/go/v4/auth"
	"github.com/golang/glog"
	"github.com/rhyth-me/backend/pkg/firebase"
)

// Client - firebase authentication client
var (
	Client      *auth.Client
	SessionName string
)

// Init - setup firebase authentication client
func Init() *auth.Client {
	ctx := context.Background()
	app := firebase.Init()

	cli, err := app.Auth(ctx)
	if err != nil {
		glog.Errorln(err)
	}

	Client = cli
	SessionName = os.Getenv("SESSION_NAME")

	return cli
}
