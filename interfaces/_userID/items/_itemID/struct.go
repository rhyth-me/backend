package item

import (
	"github.com/rhyth-me/backend/domain/model"
)

// GetRequest - fetch item info by ID.
type GetRequest struct {
	UserID string `json:"userID" param:"userID"`
	ItemID string `json:"itemID" param:"itemID"`
}

// GetResponse -
type GetResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message,omitempty"`
	Result  model.Item `json:"result,omitempty"`
}
