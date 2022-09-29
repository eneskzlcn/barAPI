package ping_pong

type PingRequest struct {
	Times int `json:"times"`
}

const (
	PongMinTimes = 1
	PongMaxTimes = 200
)

func (p PingRequest) Validate() error {
	if p.Times > PongMaxTimes || p.Times < PongMinTimes {
		return InvalidPongTimes
	}
	return nil
}
