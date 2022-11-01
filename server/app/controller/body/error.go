package body

type Error struct {
	Message string `json:"error"`
}

func NewError(s string) *Error {
	return &Error{Message: s}
}
