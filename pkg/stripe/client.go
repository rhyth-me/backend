package stripe

import (
	"os"

	"github.com/stripe/stripe-go/v72/client"
)

// Init - setup stripe client
func Init() *client.API {
	sc := &client.API{}
	sc.Init(os.Getenv("STRIPE_API_KEY"), nil)

	return sc
}
