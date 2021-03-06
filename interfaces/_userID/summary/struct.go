package summary

import "github.com/rhyth-me/backend/domain/model"

// GetRequest - fetch social profile by User ID.
type GetRequest struct {
	UserID string `json:"userID" param:"userID"`
}

// GetResponse -
type GetResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message,omitempty"`
	Result  GetResponseResult `json:"result,omitempty"`
}

// GetResponseResult -
type GetResponseResult struct {
	Profile model.SocialProfile `json:"profile,omitempty"`
}
