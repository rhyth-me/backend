package accounts

// PostStripeRequest - create new stripe connect account
type PostStripeRequest struct {
	UserID string `json:"userID" param:"userID"`
	EMail  string `json:"email" param:"email"`
}

// PostStripeResponse - return redirect URL
type PostStripeResponse struct {
	Code        int    `json:"code"`
	Message     string `json:"message,omitempty"`
	RedirectURL string `json:"redirectUrl,omitempty"`
}
