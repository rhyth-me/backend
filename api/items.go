package api

import "github.com/rhyth-me/backend/domain/model"

// PutItemsRequest - register new item
type PutItemsRequest struct {
	Details model.ItemSnippet `json:"details"`
}

// PutItemsResponse -
type PutItemsResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message,omitempty"`
	Result  model.Item `json:"result,omitempty"`
}
