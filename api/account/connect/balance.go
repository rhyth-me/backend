package connect

// GetBalanceRequest -
type GetBalanceRequest struct{}

// GetBalanceResponse -
type GetBalanceResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message,omitempty"`
	Result  map[string]int `json:"result,omitempty"`
}
