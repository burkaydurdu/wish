package wish

import "errors"

var (
	ErrInvalidEmail       = errors.New("invalid email format")
	ErrEmailAlreadyExists = errors.New("email already exists")
)
