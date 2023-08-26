package utils

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// Error implements error.
func (r Error) Error() string {
	return r.Message
}

func HandleError(message string, code int) error {
	return &Error{
		Message: message,
		Code:    code,
	}
}
