package model

// User - user
type User struct {
	ID              string `json:"id"`
	DisplayName     string `json:"displayName"`
	ProfileImageURL string `json:"profileImageUrl"`
	StatusMessage   string `json:"statusMessage"`
}
