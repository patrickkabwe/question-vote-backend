package utils

import (
	"vote-app/constants"
	"vote-app/types"
)

func LoginRequestValidate(l *types.LoginRequest) error {
	if l.Email == "" {
		return constants.ErrEmailRequired
	}

	if l.Password == "" {
		return constants.ErrPasswordRequired
	}

	return nil
}
