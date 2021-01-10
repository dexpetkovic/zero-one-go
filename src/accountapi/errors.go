package accountapi

// ResponseErr is an expected error message body format
type responseErr struct {
	ErrorMessage string `json:"error_message"`
}
