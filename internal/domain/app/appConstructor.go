package app

import (
	"WB_cloud/internal/adapters/HttpAdapter"
	"WB_cloud/internal/domain/usecases"
	"context"
	"flag"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"sync"
)

var err error
var logger *zap.Logger
var storage Storage
var httpServer HttpAdapter.HttpAdapter

func Start(ctx context.Context) {
	storageType := flag.String("storage", "in-mem",
		"defines storage type: postgres, mongo, in-mem, etc")
	flag.Parse()
	if storageType == nil {
		*storageType = "in-mem"
	}

	logger, _ = zap.NewProduction()
	storage = NewStorage(*storageType)
	accService := usecases.NewAccountService(storage)
	httpServer = HttpAdapter.New(accService, logger)

	group, gctx := errgroup.WithContext(ctx)
	group.Go(func() error { return storage.Connect(gctx) })
	group.Go(func() error { return httpServer.Start(gctx) })

	logger.Info("application is starting")

	if err = group.Wait(); err != nil {
		logger.Fatal("application start fail", zap.Error(err))
	}
}

func Stop() {
	var wg sync.WaitGroup
	ctx := context.Background()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := httpServer.Stop(ctx); err != nil {
			logger.Warn("main server shutdown error", zap.Error(err))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := storage.Close(ctx); err != nil {
			logger.Warn("error on storage closing", zap.Error(err))
		} else {
			logger.Info("storage closed gracefully")
		}
	}()

	wg.Wait()
	logger.Info("application stopped successfully")
}
