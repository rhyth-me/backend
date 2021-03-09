package stripe

import (
	"os"

	"github.com/stripe/stripe-go/v72/client"
)

// Client - stripe client
var Client *client.API

// Init - setup stripe client
func Init() *client.API {
	var apiKey string = os.Getenv("STRIPE_API_KEY")

	sc := &client.API{}
	sc.Init(apiKey, nil)

	Client = sc

	return sc
}
