package model

// HTTPResponse represents a standard HTTP response format.
type HTTPResponse struct {
	Code    int         `json:"response_code"`
	Message string      `json:"response_message"`
	Data    interface{} `json:"data,omitempty"`
}

// NewHTTPResponse creates a new HTTPResponse instance with the provided code, message, and optional data.
func NewHTTPResponse(code int, message string, data interface{}) *HTTPResponse {
	return &HTTPResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
