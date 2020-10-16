package models

import "errors"

var (
	// ErrNotFound is an error when query dont retrieve any result
	ErrNotFound = errors.New("resquesed item is not found")
)
