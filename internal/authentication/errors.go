package authentication

// ErrClientNotFound error representing that authData wasn't found
type ErrClientNotFound struct {
	err error
}

func NewClientNotFoundError(err error) ErrClientNotFound {
	return ErrClientNotFound{err}
}

// Error returns error's description
func (e ErrClientNotFound) Error() string {
	msg := "error while getting a client"

	if e.err != nil {
		msg += ": " + e.err.Error()
	}

	return msg
}

// Unwrap returns a wrapped error
func (e ErrClientNotFound) Unwrap() error {
	return e.err
}

// ErrClientNotFound error representing that a password is incorrect
type ErrInvalidPassword struct {
	err error
}

func NewInvalidPasswordError(err error) ErrInvalidPassword {
	return ErrInvalidPassword{err}
}

// Error returns error's description
func (e ErrInvalidPassword) Error() string {
	msg :=  "error while getting a client"

	if e.err != nil {
		msg += ": " + e.err.Error()
	}

	return msg
}

// Unwrap returns a wrapped error
func (e ErrInvalidPassword) Unwrap() error {
	return e.err
}
