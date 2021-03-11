package account

import "github.com/rhyth-me/backend/domain/model"

// GetProfileRequest - fetch auth user's profile
type GetProfileRequest struct{}

// GetProfileResponse - return auth user's profile
type GetProfileResponse struct {
	Code    int                 `json:"code"`
	Message string              `json:"message,omitempty"`
	Result  model.SocialProfile `json:"result,omitempty"`
}

// PatchProfileRequest - update auth user's profile
type PatchProfileRequest struct {
	Profile model.SocialProfile `json:"profile" validate:"required"`
}

// PatchProfileResponse - return auth user's profile
type PatchProfileResponse struct {
	Code    int                 `json:"code"`
	Message string              `json:"message,omitempty"`
	Result  model.SocialProfile `json:"result,omitempty"`
}
