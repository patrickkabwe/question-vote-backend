package constants

import "errors"

var (
	ErrEmailRequired      = errors.New("email is required")
	ErrPasswordRequired   = errors.New("password is required")
	ErrAllFieldsRequired  = errors.New("all fields are required")
	ErrEmailInvalid       = errors.New("email is invalid")
	ErrInValidCredentials = errors.New("invalid credentials")
)
