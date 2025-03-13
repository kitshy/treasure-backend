package main

import (
	"context"
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/kitshy/treasure-backend/common/opio"
)

func main() {
	log.SetDefault(log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stderr, log.LevelInfo, true)))
	app := NewCli("GitCommit", "GitData")
	ctx := opio.WithInterruptBlocker(context.Background())
	if err := app.RunContext(ctx, os.Args); err != nil {
		log.Error("application error", "err", err)
		os.Exit(1)
	}
}
