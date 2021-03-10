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

// DeleteCardsRequest - Delete the card with the specified number.
type DeleteCardsRequest struct {
	CardID int `json:"cardId" validate:"min=1,max=5"`
}

// DeleteCardsResponse - Return the result of deletion.
type DeleteCardsResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}
