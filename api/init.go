package api

import (
  "time"
)

// JSON is a standard form of response data
type JSON struct {
	Status    string      `json:"status"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	CreatedAt string      `json:"createdAt"`
}

const (
)

// NewJSON creates a new JSON object
func NewJSON(status, message string, data interface{}) *JSON {
	return &JSON{
		Status:  status,
		Message: message,
		Data:    data,
    CreatedAt: time.Now().String(),
	}
}
