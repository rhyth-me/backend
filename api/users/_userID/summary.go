package users

import "github.com/rhyth-me/backend/domain/model"

// GetSummaryRequest - Fetch user summary by User ID.
type GetSummaryRequest struct {
	UserID string `json:"userID" param:"userID"`
}

// GetSummaryResponse - Return user summary by User ID.
type GetSummaryResponse struct {
	Code    int                      `json:"code"`
	Message string                   `json:"message,omitempty"`
	Result  GetSummaryResponseResult `json:"result,omitempty"`
}

// GetSummaryResponseResult - Structure of user summary.
type GetSummaryResponseResult struct {
	Profile model.SocialProfile `json:"profile,omitempty"`
}
