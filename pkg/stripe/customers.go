package stripe

import (
	"errors"

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

	c, err := Client.Customers.New(params)
	if err != nil {
		return nil, errors.New("Failed to create customer")
	}

	return c, nil
}
