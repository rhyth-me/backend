package account

import "github.com/rhyth-me/backend/domain/model"

// GetCardsRequest - Fetch the list of card information.
type GetCardsRequest struct{}

// GetCardsResponse - Return the list of card information.
type GetCardsResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message,omitempty"`
	Result  []model.Card `json:"result,omitempty"`
}

// PutCardsRequest - Create a new card in Stripe
type PutCardsRequest struct {
	Token string `json:"token"`
}

// PutCardsResponse - Return the information of the created card.
type PutCardsResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message,omitempty"`
	Result  model.Card `json:"result,omitempty"`
}
