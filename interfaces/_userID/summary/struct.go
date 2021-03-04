package summary

// GetRequest - fetch social profile by User ID.
type GetRequest struct {
	UserID string `json:"userID" param:"userID"`
}

// GetResponse -
type GetResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}
