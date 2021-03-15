package model

// User - users collection structure.
type User struct {
	UID     string        `firestore:"uid" json:"-"`
	Google  Google        `firestore:"google" json:"-"`
	Profile SocialProfile `firestore:"profile" json:"profile,omitempty"`
	Payment Payment       `firestore:"payment" json:"-"`
}

// Field Structures

// Google - The user's Google account.
type Google struct {
	ID    string `firestore:"id" json:"-"`
	Email string `firestore:"email"  json:"-"`
}

// Access - user browser env
type Access struct {
	IPAddress string `firestore:"ipAddress" json:"ipAddress"`
	UserAgent string `firestore:"userAgent" json:"userAgent"`
}

// SocialProfile - details of the user as displayed on the site.
type SocialProfile struct {
	ScreenName    string `firestore:"screenName" json:"screenName" validate:"omitempty,min=3,max=25"`
	DisplayName   string `firestore:"displayName" json:"displayName" validate:"omitempty,min=1,max=20"`
	StatusMessage string `firestore:"statusMessage" json:"statusMessage" validate:"omitempty,min=0,max=150"`
	ImageHash     string `firestore:"imageHash" json:"imageHash" validate:"omitempty,len=64"`
}

// PayoutSettings - information for withdrawal.
type PayoutSettings struct {
	StripeID string `firestore:"stripeId"`
}
