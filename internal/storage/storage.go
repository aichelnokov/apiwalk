package storage

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserExists   = errors.New("user already exists")
	ErrWalkNotFound = errors.New("walk not found")
)