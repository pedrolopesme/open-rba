package collector

import "errors"

var (
	ErrEmptyUserId = errors.New("no user ID was provided")
	ErrEmptyIP     = errors.New("no IP address was provided")
)
