package connect

// GetLoginRequest -
type GetLoginRequest struct{}

// GetLoginResponse -
type GetLoginResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Result  string `json:"result,omitempty"`
}
