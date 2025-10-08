package middleware

import "errors"

var (
	ErrFWNotFound = errors.New("no supported framework modules found in go.mod")
)
