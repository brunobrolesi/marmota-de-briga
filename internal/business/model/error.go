package model

import "errors"

var (
	ErrClientLimitExceeded = errors.New("client limit exceeded")
	ErrInternalServerError = errors.New("internal server error")
)
