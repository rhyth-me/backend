package users

import "github.com/rhyth-me/backend/domain/model"

// GetItemsRequest - fetch item list by User ID.
type GetItemsRequest struct {
	UserID string `json:"userID" param:"userID"`
}

// GetItemsResponse -
type GetItemsResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message,omitempty"`
	Result  []model.Item `json:"result,omitempty"`
}
