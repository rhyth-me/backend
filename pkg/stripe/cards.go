package stripe

import (
	"errors"
	"strings"

	"github.com/stripe/stripe-go/v72"
)

var customerPrefix string = "cus_"

// CreateCard - Create a new card on Stripe.
func CreateCard(CustomerID string, token string) (*stripe.Card, error) {
	if !strings.HasPrefix(CustomerID, customerPrefix) {
		return nil, errors.New("Invalid Customer ID")
	}

	params := &stripe.CardParams{
		Customer: stripe.String(CustomerID),
		Token:    stripe.String(token),
	}

	return Client.Cards.New(params)
}
