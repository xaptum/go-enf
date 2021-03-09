package enf

import (
	"fmt"
	"strings"
)

// ErrorResponse represents the error response from the API.
type EnfApiError struct {
	StatusCode   int     `json:"-"`
	ErrorMessage *string `json:"-"`
	CodeError    *struct {
		Code string `json:"code"`
		Text string `json:"text"`
	} `json:"error"`
	ReasonError *struct {
		Reason string `json:"reason"`
	} `json:"xiam_error"`
}

func (e *EnfApiError) Error() string {
	var msg string

	if nil != e.CodeError {
		msg = fmt.Sprintf("%v: %v", strings.ToUpper(e.CodeError.Code), e.CodeError.Text)
	} else if nil != e.ReasonError {
		msg = fmt.Sprintf("%v", e.ReasonError.Reason)
	} else if nil != e.ErrorMessage {
		msg = fmt.Sprintf("%v", e.ErrorMessage)
	} else {
		msg = "UNKNOWN_ERROR: server did not respond with properly formatted error message."
	}

	return msg
}
