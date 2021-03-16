package account

// GetLoginRequest - Create a new session cookie.
type GetLoginRequest struct {
	IDtoken string `query:"idToken" json:"idToken"`
}

// GetLoginResponse - Set a new session cookie.
type GetLoginResponse struct {
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
