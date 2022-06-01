package command

import (
	"fmt"
	"os"
	"context"
)

const (
	exitOk    = 0
	exitError = 1
)

func Run() {
	fmt.Println("start")
	os.Exit(run(context.Background())) // 空のコンテキスト生成	
}

func run(ctx context.Context) int {
	return
}