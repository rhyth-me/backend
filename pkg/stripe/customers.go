package stripe

import (
	"github.com/rhyth-me/backend/domain/model"
	"github.com/stripe/stripe-go/v72"
)

// CreateCustomer - Create new customer on Stripe.
func CreateCustomer(User *model.User) (*stripe.Customer, error) {

	params := &stripe.CustomerParams{
		Params: stripe.Params{
			Metadata: map[string]string{
				"GoogleID":    User.Google.ID,
				"Email":       User.Google.Email,
				"FirebaseUID": User.UID,
			},
		},
		Email: &User.Google.Email,
	}

	return Client.Customers.New(params)
}
