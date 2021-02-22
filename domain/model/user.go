package model

// User - user
type User struct {
	ID              int    `json:"id"`
	ScreenName      string `json:"screenName"`
	DisplayName     string `json:"displayName"`
	ProfileImageURL string `json:"profileImageUrl"`
	StatusMessage   string `json:"statusMessage"`
}
