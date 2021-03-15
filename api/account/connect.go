package account

// GetConnectRequest - Fetch information about authed user's connect account.
type GetConnectRequest struct{}

// GetConnectResponse - Return the information.
type GetConnectResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Result  string `json:"result,omitempty"`
}

// PutConnectRequest - Submit the consent information to stripe & create a new connect account.
type PutConnectRequest struct{}

// PutConnectResponse - Return the result.
type PutConnectResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}
