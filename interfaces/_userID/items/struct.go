package items

import "github.com/rhyth-me/backend/domain/model"

// GetRequest - fetch item list by User ID.
type GetRequest struct {
	UserID string `json:"userID" param:"userID"`
}

// GetResponse -
type GetResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message,omitempty"`
	Result  []model.Item `json:"result,omitempty"`
}
