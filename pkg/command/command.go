package command

import (
	"fmt"
	"os"
	"context"

	"exercise-go-api/pkg/version"

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
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to setup logger: %s\n", err)
		return exitError
	}
	defer logger.Sync()
	logger = logger.With(zap.String("version", version.Version))
	return exitOk
}