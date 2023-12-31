package errorops

type Error struct {
	Code        int      `json:"code,omitempty"`
	Description string   `json:"description,omitempty"`
	Message     []string `json:"message,omitempty"`
	Value       any      `json:"value,omitempty"`
}

func NewError(code int, description string, value any, message ...string) *Error {
	return &Error{
		Code:        code,
		Value:       value,
		Description: description,
		Message:     message,
	}
}

func (e *Error) Error() string {
	return e.Description
}
