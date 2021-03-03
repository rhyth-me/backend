package model

// User - users collection structure.
type User struct {
	UID     string         `firestore:"uid" json:"-"`
	Profile SocialProfile  `firestore:"profile" json:"profile,omitempty"`
	Payout  PayoutSettings `firestore:"payout" json:"-"`
}

// Field Structures

// SocialProfile - details of the user as displayed on the site.
type SocialProfile struct {
	ID               string `firestore:"id" json:"id" validate:"gt=0,max=20"`
	DisplayName      string `firestore:"displayName" json:"displayName" validate:"gt=0,max=20"`
	ProfileImagePath string `firestore:"profileImagePath" json:"profileImagePath" validate:"required"`
	StatusMessage    string `firestore:"statusMessage" json:"statusMessage" validate:"min=0,max=150"`
}

// PayoutSettings - information for withdrawal.
type PayoutSettings struct {
	StripeID string `firestore:"stripeId"`
}
