package user

import "errors"

var (
	ErrUserExists   = errors.New("user already exists")
	ErrInvalidName  = errors.New("invalid user name")
	ErrUserNotFound = errors.New("user not found")
)
