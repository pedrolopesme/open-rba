package redisclient

import "errors"

var (
	ErrClientNotConnect = errors.New("Not connected to redis")
)
