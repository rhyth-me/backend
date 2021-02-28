package items

import (
	"github.com/rhyth-me/backend/domain/model"
)

// GetItemInfoRequest - fetch item info by ID.
type GetItemInfoRequest struct {
	UserID string `json:"userID" param:"userID"`
	ItemID string `json:"itemID" param:"itemID"`
}

// GetItemInfoResponse -
type GetItemInfoResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message,omitempty"`
	Result  model.Item `json:"result,omitempty"`
}
