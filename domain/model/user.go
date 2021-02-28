package model

// User - user
type User struct {
	ID              string       `firestore:"id" json:"id"`
	DisplayName     string       `firestore:"displayName" json:"displayName"`
	ProfileImageURL string       `firestore:"profileImageUrl" json:"profileImageUrl"`
	StatusMessage   string       `firestore:"statusMessage" json:"statusMessage"`
	Accounts        UserAccounts `firestore:"accounts" json:"accounts"`
}

// UserAccounts - stripe
type UserAccounts struct {
	Stripe string `firestore:"stripe" json:"stripe"`
}
