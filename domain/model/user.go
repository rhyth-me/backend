package model

// User - user
type SocialProfile struct {
	ID              string `firestore:"id" json:"id"`
	DisplayName     string `firestore:"displayName" json:"displayName"`
	ProfileImageURL string `firestore:"profileImageUrl" json:"profileImageUrl"`
	StatusMessage   string `firestore:"statusMessage" json:"statusMessage"`
}

// AuthUser -
type AuthUser struct {
	UID string `json:"uid"`
}

// UserAccounts - stripe
type UserAccounts struct {
	Stripe string `firestore:"stripe" json:"stripe"`
}
