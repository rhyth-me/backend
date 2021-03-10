package stripe

import (
	"errors"
	"strings"

	"github.com/stripe/stripe-go/v72"
)

// CreateCard - Create a new card on Stripe.
func CreateCard(CustomerID string, token string) (*stripe.Card, error) {
	if !strings.HasPrefix(CustomerID, CustomerPrefix) {
		return nil, errors.New("Invalid Customer ID")
	}

	params := &stripe.CardParams{
		Customer: stripe.String(CustomerID),
		Token:    stripe.String(token),
	}

	c, err := Client.Cards.New(params)
	if err != nil {
		return nil, errors.New("Failed to create card")
	}

	return c, nil
}

// DeleteCard - Delete the specified card.
func DeleteCard(CustomerID string, CardID string) (*stripe.Card, error) {
	if !strings.HasPrefix(CustomerID, CustomerPrefix) {
		return nil, errors.New("Invalid Customer ID")
	}

	if !strings.HasPrefix(CardID, CardPrefix) {
		return nil, errors.New("Invalid Card ID")
	}

	params := &stripe.CardParams{
		Customer: stripe.String(CustomerID),
	}

	c, err := Client.Cards.Del(CardID, params)
	if err != nil {
		return nil, errors.New("Failed to delete card")
	}

	return c, nil
}
