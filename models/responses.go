// models/responses.go
package models

type DetailedErrorResponse struct {
	Status       string      `json:"status"`
	StatusCode   int         `json:"statusCode"`
	Message      interface{} `json:"message"` // Use interface{} to accept any data type
	ShowNextFlow bool        `json:"showNextFlow"`
	RefId        interface{} `json:"refId"`
	Payload      interface{} `json:"payload"`
	File         interface{} `json:"file"`
}

// ErrorResponse represents a JSON structure for API error responses.
type ErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}
