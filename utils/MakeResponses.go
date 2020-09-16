package utils

import (
	"github.com/scafol/KP-Backend/responsegraph"
)

// MakeResponse - create a request response
func MakeResponse(Status string, Message string, Data interface{}) responsegraph.ResponseGeneric {
	return responsegraph.ResponseGeneric{
		Status:  Status,
		Message: Message,
		Data:    Data,
	}
}
