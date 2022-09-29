package ping_pong

type Service struct {
	logger Logger
}

func NewService(logger Logger) *Service {
	if logger == nil {
		return nil
	}
	return &Service{logger: logger}
}
func (s *Service) Ping(request PingRequest) (PongResponse, error) {
	s.logger.Debugf("Creating pongs amount of %d", request.Times)
	if err := request.Validate(); err != nil {
		return PongResponse{}, err
	}
	pongs := s.createPongs(request.Times)

	return PongResponse{Pongs: pongs}, nil
}
func (s *Service) createPongs(times int) []string {
	pongs := make([]string, 0)
	for i := 0; i < times; i++ {
		pongs = append(pongs, PONG)
	}
	return pongs
}
