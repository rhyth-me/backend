package account

// GetConnectRequest - Fetch information about authed user's connect account.
type GetConnectRequest struct{}

// GetConnectResponse - Return the information.
type GetConnectResponse struct {
	Code    int                      `json:"code"`
	Message string                   `json:"message,omitempty"`
	Result  GetConnectResponseResult `json:"result,omitempty"`
}

type GetConnectResponseResult struct {
	Exist  bool `json:"exist"`
	Status int  `json:"status"` // 0:need info, 1:active, -1:banned
}

// PostConnectRequest - Submit the consent information to stripe & create a new connect account.
type PostConnectRequest struct{}

// PostConnectResponse - Return the result.
type PostConnectResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}

// PatchConnectRequest - Update status.
type PatchConnectRequest struct{}

// PatchConnectResponse - Return status.
type PatchConnectResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Result  int    `json:"result,omitempty"`
}
