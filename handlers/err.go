package handlers

// Error returned by the API
type Error struct {
	// Error text
	Error string `json:"error"`
}

// NewError creates a new error from the provided standard error
func NewError(err error) Error {
	if err == nil {
		return Error{}
	}

	return Error{
		Error: err.Error(),
	}
}
