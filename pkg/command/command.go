package command

import (
	"fmt"
	"os"
	"context"
	"net"

	"exercise-go-api/pkg/version"
	"exercise-go-api/pkg/config"

	"go.uber.org/zap"
)

const (
	exitOk    = 0
	exitError = 1
)

func Run() {
	os.Exit(run(context.Background())) // 空のコンテキスト生成	
}

func run(ctx context.Context) int {
	// init logger
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to setup logger: %s\n", err)
		return exitError
	}
	defer logger.Sync()
	logger = logger.With(zap.String("version", version.Version))

	// load config
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		logger.Error("failed to load config", zap.Error(err))
		return exitError
	}

	// init listener
	listener, err := net.Listen("tcp", cfg.Address())
	if err != nil {
		logger.Error("failed to listen port", zap.Int("port", cfg.Port), zap.Error(err))
		return exitError
	}
	logger.Info("server start listening", zap.Int("port", cfg.Port)) // ClIにログを書き出す

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	return exitOk

}