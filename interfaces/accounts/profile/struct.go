package profile

import "github.com/rhyth-me/backend/domain/model"

// GetRequest - fetch auth user's profile
type GetRequest struct{}

// GetResponse - return auth user's profile
type GetResponse struct {
	Code    int                 `json:"code"`
	Message string              `json:"message,omitempty"`
	Result  model.SocialProfile `json:"result,omitempty"`
}

// PutRequest - fetch auth user's profile
type PutRequest struct {
	Profile model.SocialProfile `json:"profile"`
}

// PutResponse - return auth user's profile
type PutResponse struct {
	Code    int                 `json:"code"`
	Message string              `json:"message,omitempty"`
	Result  model.SocialProfile `json:"result,omitempty"`
}
