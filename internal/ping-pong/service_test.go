package ping_pong_test

import (
	"errors"
	"github.com/eneskzlcn/ping-pong/internal/mocks"
	ping_pong "github.com/eneskzlcn/ping-pong/internal/ping-pong"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createMockLogger(t *testing.T) *mocks.MockLogger {
	return mocks.NewMockLogger(gomock.NewController(t))
}
func TestNewService(t *testing.T) {
	mockLogger := createMockLogger(t)
	t.Run("given empty logger then it should return nil", func(t *testing.T) {
		service := ping_pong.NewService(nil)
		assert.Nil(t, service)
	})
	t.Run("given valid arguments then it should return service", func(t *testing.T) {
		service := ping_pong.NewService(mockLogger)
		assert.NotNil(t, service)
	})
}
func createPongsByGivenTime(times int) []string {
	pongs := make([]string, 0)
	for i := 0; i < times; i++ {
		pongs = append(pongs, ping_pong.PONG)
	}
	return pongs
}
func TestService_Ping(t *testing.T) {
	mockLogger := createMockLogger(t)
	service := ping_pong.NewService(mockLogger)
	mockLogger.EXPECT().Debugf(gomock.Any(), gomock.Any()).AnyTimes()
	t.Run("given invalid pong times in request then it should return InvalidPongTimes error", func(t *testing.T) {
		invalidPingRequest := ping_pong.PingRequest{Times: -1}
		_, err := service.Ping(invalidPingRequest)
		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, ping_pong.InvalidPongTimes))
	})
	t.Run("given valid pong times in request then it should return pong response contains pongs as given times", func(t *testing.T) {
		givenPingRequest := ping_pong.PingRequest{Times: 20}
		expectedPongs := createPongsByGivenTime(givenPingRequest.Times)
		expectedPongResponse := ping_pong.PongResponse{Pongs: expectedPongs}

		pongResponse, err := service.Ping(givenPingRequest)
		assert.Nil(t, err)
		assert.Equal(t, expectedPongResponse, pongResponse)
	})
}
