package model

// User - users collection structure.
type User struct {
	UID     string         `firestore:"uid" json:"-"`
	Google  Google         `firestore:"google" json:"-"`
	Profile SocialProfile  `firestore:"profile" json:"profile,omitempty"`
	Payout  PayoutSettings `firestore:"payout" json:"-"`
}

// Field Structures

// Google - The user's Google account.
type Google struct {
	ID    string `firestore:"id" json:"-"`
	Email string `firestore:"email"  json:"-"`
}

// SocialProfile - details of the user as displayed on the site.
type SocialProfile struct {
	ScreenName       string `firestore:"screenName" json:"screenName" validate:"gt=0,max=25"`
	DisplayName      string `firestore:"displayName" json:"displayName" validate:"gt=0,max=20"`
	ProfileImagePath string `firestore:"profileImagePath" json:"profileImagePath" validate:"required"`
	StatusMessage    string `firestore:"statusMessage" json:"statusMessage" validate:"min=0,max=150"`
}

// PayoutSettings - information for withdrawal.
type PayoutSettings struct {
	StripeID string `firestore:"stripeId"`
}
