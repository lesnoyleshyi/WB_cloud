package main

import (
	"WB_cloud/internal/domain/app"
	"context"
	"os/signal"
	"syscall"
)

func main() {
	baseCtx := context.Background()
	ctx, cancel := signal.NotifyContext(baseCtx,
		syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	defer cancel()

	go app.Start(ctx)

	<-ctx.Done()
	app.Stop()
}
