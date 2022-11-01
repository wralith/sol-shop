package body

type ErrorMessage struct {
	Message string `json:"error"`
}

func NewErrorMessage(s string) *ErrorMessage {
	return &ErrorMessage{Message: s}
}
