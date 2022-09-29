package ping_pong

import "errors"

const (
	PONG = "pong"
)

var (
	InvalidPongTimes = errors.New("invalid pong times")
)
