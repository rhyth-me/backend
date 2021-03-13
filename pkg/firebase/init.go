package firebase

import (
	"context"
	"os"

	firebase "firebase.google.com/go/v4"
	"github.com/golang/glog"
	"google.golang.org/api/option"
)

// Init - setup firebase client
func Init() *firebase.App {
	config := &firebase.Config{
		StorageBucket: "score-market.appspot.com",
	}

	opt := option.WithCredentialsJSON([]byte(os.Getenv("SERVICE_ACCOUNT_KEY")))

	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		glog.Errorln("Error")
	}
	return app
}
