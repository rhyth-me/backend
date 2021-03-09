package model

// Payment - Information about the user's payment.
type Payment struct {
	CustomerID string `firestore:"customerId" json:"-"` // prefix - `cus_`
	ConnectID  string `firestore:"connectId" json:"-"`  // prefix - `acct_`
	Cards      []Card `firestore:"cards" json:"-"`
}

// Card - Information about the card.
type Card struct {
	ID       string `firestore:"id" json:"-"` // prefix - `card_`
	Brand    string `firestore:"brand" json:"brand"`
	Country  string `firestore:"country" json:"country"`
	Last4    string `firestore:"last4" json:"last4"`
	ExpYear  string `firestore:"expYear" json:"expYear"`
	ExpMonth string `firestore:"expMonth" json:"expMonth"`
}
