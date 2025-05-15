package helpers

import "errors"

var (
	ErrNotFound = errors.New("not.found")
	ErrUnkown   = errors.New("unknown.error")
)
