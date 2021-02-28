package accounts

// PostCreateAccountRequest - create new stripe connect account
type PostCreateAccountRequest struct {
	UserID string `json:"userID" param:"userID"`
	EMail  string `json:"email" param:"email"`
}

// PostCreateAccountResponse - return redirect URL
type PostCreateAccountResponse struct {
	Code        int    `json:"code"`
	Message     string `json:"message,omitempty"`
	RedirectURL string `json:"redirectUrl,omitempty"`
}
