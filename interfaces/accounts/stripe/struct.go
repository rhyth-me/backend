package stripe

// PostRequest - create new stripe connect account
type PostRequest struct {
	EMail string `json:"email" param:"email"`
}

// PostResponse - return redirect URL
type PostResponse struct {
	Code        int    `json:"code"`
	Message     string `json:"message,omitempty"`
	RedirectURL string `json:"redirectUrl,omitempty"`
}
