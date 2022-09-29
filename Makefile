build:
	go build -o bin/ping-pong ./cmd/ping-pong

run:
	./bin/ping-pong

start:
	go build -o bin/ping-pong ./cmd/ping-pong && ./bin/ping-pong

dockerize:
	docker build -t eneskzlcn/ping-pong:latest .

dockerun:
	docker run -p 4200:4200 eneskzlcn/ping-pong:latest
clean:
	rm -rf bin

unit-tests:
	go test -v ./internal/ping-pong

generate-mocks:
	mockgen -destination=internal/mocks/mock_logger.go -package mocks github.com/eneskzlcn/ping-pong/internal/ping-pong Logger
	mockgen -destination=internal/mocks/mock_ping-pong_service.go -package mocks github.com/eneskzlcn/ping-pong/internal/ping-pong PingPongService