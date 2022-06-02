package command

import (
	"fmt"
	"os"
	"context"
	"net"
	"os/signal"
	"syscall"

	"exercise-go-api/pkg/version"
	"exercise-go-api/pkg/config"
	"exercise-go-api/pkg/server"
	"exercise-go-api/pkg/handler"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
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
	defer logger.Sync() //関数を抜ける際にログをフラッシュしてバッファに残っているログレコードを書き込む
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
	logger.Info("server start listening", zap.Int("port", cfg.Port)) // ログを書き出す

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// init mongo db
	// logger.Info("connect to mongo db", zap.String("url", cfg.DB.URL), zap.String("source", cfg.DB.Source))
	// opts := &options.ClientOptions{}
	// if cfg.DB.Source == "external" {
	// 	opts = options.Client().SetAuth(options.Credential{AuthMechanism: "MONGODB-AWS", AuthSource: "$external"})
	// }

	// mongoClient, err := mongo.NewClient(options.Client().ApplyURI(cfg.DB.URL), opts)
	// if err != nil {
	// 	logger.Error("failed to create mongo db client", zap.Error(err), zap.String("uri", cfg.DB.URL))
	// 	return exitError
	// }

	// mongoCtx, mongoCancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer mongoCancel()

	// if err := mongoClient.Connect(mongoCtx); err != nil {
	// 	logger.Error("failed to connect to mongo db", zap.Error(err))
	// 	return exitError
	// }

	// if err := mongoClient.Ping(mongoCtx, readpref.Primary()); err != nil {
	// 	logger.Error("failed to ping mongo db", zap.Error(err))
	// 	return exitError
	// }
	// mongoDB := mongoClient.Database(cfg.DB.Database)

	// get repositories
	// repositories, err := persistence.NewRepositories(mongoDB)
	// if err != nil {
	// 	logger.Error("failed to new repositories", zap.Error(err))
	// 	return exitError
	// }

	// init to start http server
	registry := handler.NewHandler(logger, repositories, version.Version)
	httpServer := server.NewServer(registry, &server.Config{Log: logger})
	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(func() error {
		return httpServer.Serve(listener)
	})

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, os.Interrupt)
	select {
	case <-sigCh:
	case <-ctx.Done():
	}

	if err := httpServer.GracefulShutdown(ctx); err != nil {
		return exitError
	}

	cancel()
	if err := wg.Wait(); err != nil {
		return exitError
	}


	return exitOk

}