package ping_pong_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/eneskzlcn/ping-pong/internal/mocks"
	ping_pong "github.com/eneskzlcn/ping-pong/internal/ping-pong"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHandler(t *testing.T) {
	mockLogger, mockService := createHandlerMocks(t)
	t.Run("given nil logger then it should return nil", func(t *testing.T) {
		handler := ping_pong.NewHandler(nil, mockService)
		assert.Nil(t, handler)
	})
	t.Run("given nil service then it should return nil", func(t *testing.T) {
		handler := ping_pong.NewHandler(mockLogger, nil)
		assert.Nil(t, handler)
	})
	t.Run("given valid arguments then it should return handler", func(t *testing.T) {
		handler := ping_pong.NewHandler(mockLogger, mockService)
		assert.NotNil(t, handler)
	})
}
func createHandlerMocks(t *testing.T) (*mocks.MockLogger, *mocks.MockPingPongService) {
	ctrl := gomock.NewController(t)
	return mocks.NewMockLogger(ctrl), mocks.NewMockPingPongService(ctrl)
}
func makeTestRequestWithBody(app *fiber.App, method string, route string, body interface{}) (*http.Response, error) {
	req := createRequestWithBody(method, route, body)
	return app.Test(req)
}
func createRequestWithBody(method, route string, body interface{}) *http.Request {
	bodyAsByte, _ := json.Marshal(body)
	req := httptest.NewRequest(method, route, bytes.NewReader(bodyAsByte))
	req.Header.Add("Content-type", "application/json")
	return req
}
func assertJSONBodyEqual(t *testing.T, body io.Reader, expected interface{}) {
	var actualBody interface{}
	_ = json.NewDecoder(body).Decode(&actualBody)

	expectedJSON, _ := json.Marshal(expected)

	var expectedBody interface{}
	_ = json.Unmarshal(expectedJSON, &expectedBody)
	assert.Equal(t, expectedBody, actualBody)
}
func TestHandler_Ping(t *testing.T) {
	mockLogger, mockService := createHandlerMocks(t)
	handler := ping_pong.NewHandler(mockLogger, mockService)
	mockLogger.EXPECT().Debugf(gomock.Any(), gomock.Any()).AnyTimes()
	app := setupFiberAppForHandler(handler)
	t.Run("given invalid ping request then it should return status bad request", func(t *testing.T) {
		pingRequest := "invalid ping request"
		resp, err := makeTestRequestWithBody(app, fiber.MethodPost, "/ping", pingRequest)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("given valid ping request but error happened in service then it should return internal server error", func(t *testing.T) {
		pingRequest := ping_pong.PingRequest{Times: 3}
		mockService.EXPECT().Ping(pingRequest).Return(ping_pong.PongResponse{}, errors.New("any"))
		resp, err := makeTestRequestWithBody(app, fiber.MethodPost, "/ping", pingRequest)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})
	t.Run("given valid ping request then it should return status ok with expected pongs response", func(t *testing.T) {
		pingRequest := ping_pong.PingRequest{Times: 3}
		expectedPongResponse := ping_pong.PongResponse{Pongs: []string{ping_pong.PONG, ping_pong.PONG, ping_pong.PONG}}
		mockService.EXPECT().Ping(pingRequest).Return(expectedPongResponse, nil)
		resp, err := makeTestRequestWithBody(app, fiber.MethodPost, "/ping", pingRequest)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		assertJSONBodyEqual(t, resp.Body, expectedPongResponse)
	})
}
func setupFiberAppForHandler(handler *ping_pong.Handler) *fiber.App {
	app := fiber.New()
	app.Post("/ping", handler.Ping)
	return app
}

func TestHandler_RegisterRoutes(t *testing.T) {
	mockLogger, mockService := createHandlerMocks(t)
	handler := ping_pong.NewHandler(mockLogger, mockService)
	app := fiber.New()
	handler.RegisterRoutes(app)

	mockLogger.EXPECT().Debugf(gomock.Any(), gomock.Any()).Times(1)
	req := httptest.NewRequest(fiber.MethodPost, "/ping", nil)
	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.NotEqual(t, fiber.StatusNotFound, resp.Status)

}
