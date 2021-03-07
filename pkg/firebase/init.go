package firebase

import (
	"context"
	"os"

	firebase "firebase.google.com/go"
	"github.com/golang/glog"
	"google.golang.org/api/option"
)

// Init - setup firebase client
func Init() *firebase.App {
	opt := option.WithCredentialsJSON([]byte(os.Getenv("SERVICE_ACCOUNT_KEY")))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		glog.Errorln("Error")
	}
	return app
}
