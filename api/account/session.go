package account

// PostLoginRequest - Create a new session cookie.
type PostLoginRequest struct {
	IDtoken string `query:"idToken" json:"idToken"`
}

// PostLoginResponse - Set a new session cookie.
type PostLoginResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

// GetLogoutRequest - Revoke current session token.
type GetLogoutRequest struct{}

// GetLogoutResponse - Return the result.
type GetLogoutResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}
