package main

import (
	"fmt"
	"github.com/eneskzlcn/ping-pong/internal/config"
	ping_pong "github.com/eneskzlcn/ping-pong/internal/ping-pong"
	"github.com/eneskzlcn/ping-pong/logger"
	"github.com/eneskzlcn/ping-pong/server"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
func getEnv(key string, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}
func run() error {
	env := getEnv("DEPLOY_ENV", "local")

	configs, err := config.LoadConfig[config.Config](".dev/", env, "yaml")

	if err != nil {
		return err
	}

	zapLogger, err := logger.NewZapLoggerForEnv(env, 0)
	if err != nil {
		return err
	}
	pingPongService := ping_pong.NewService(zapLogger)
	pingPongHandler := ping_pong.NewHandler(zapLogger, pingPongService)

	server := server.New([]server.Handler{
		pingPongHandler,
	}, configs.Server, zapLogger)

	return server.Start()
}
