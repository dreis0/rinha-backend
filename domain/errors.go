package domain

import "errors"

var (
	ErrNotFound   = errors.New("customer not found")
	ErrNotAllowed = errors.New("transaction not allowed, limit exceeded")
)
