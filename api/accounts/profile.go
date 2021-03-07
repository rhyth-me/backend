package accounts

import "github.com/rhyth-me/backend/domain/model"

// GetProfileRequest - fetch auth user's profile
type GetProfileRequest struct{}

// GetProfileResponse - return auth user's profile
type GetProfileResponse struct {
	Code    int                 `json:"code"`
	Message string              `json:"message,omitempty"`
	Result  model.SocialProfile `json:"result,omitempty"`
}

// PutProfileRequest - fetch auth user's profile
type PutProfileRequest struct {
	Profile model.SocialProfile `json:"profile"`
}

// PutProfileResponse - return auth user's profile
type PutProfileResponse struct {
	Code    int                 `json:"code"`
	Message string              `json:"message,omitempty"`
	Result  model.SocialProfile `json:"result,omitempty"`
}
