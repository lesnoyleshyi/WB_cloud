package HttpAdapter

import (
	ports "WB_cloud/internal/ports/input"
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"time"
)

const HttpAddr = `:8080`
const gracefulShutdownDelaySec = 30

type HttpAdapter struct {
	server   *http.Server
	accounts ports.AccountService
	logger   *zap.Logger
}

func New(accService ports.AccountService, logger *zap.Logger) HttpAdapter {
	var adapter HttpAdapter

	adapter.accounts = accService
	adapter.logger = logger
	server := http.Server{
		Addr:    HttpAddr,
		Handler: adapter.routes(),
	}
	adapter.server = &server

	return adapter
}

func (a HttpAdapter) Start(ctx context.Context) error {
	srvErrChan := make(chan error)

	go func() {
		if err := a.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			srvErrChan <- fmt.Errorf("couldn't start server: %w", err)
		}
		srvErrChan <- nil
	}()

	select {
	case <-ctx.Done():
		return nil
	case err := <-srvErrChan:
		return err
	}
}

func (a HttpAdapter) Stop(ctx context.Context) error {
	if a.server == nil {
		a.logger.Info("main server wasn't initialised, stop() is no-op")
		return nil
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second*gracefulShutdownDelaySec)
	defer cancel()

	err := a.server.Shutdown(timeoutCtx)
	if err != nil && errors.Is(err, http.ErrServerClosed) {
		return err
	}

	a.logger.Info("main server stopped gracefully")

	return nil
}

func (a HttpAdapter) routes() http.Handler {
	r := chi.NewRouter()

	r.Mount("/", a.routeAccounts())

	return r
}

func (a HttpAdapter) respondSuccess(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if _, err := fmt.Fprintf(w, "{\"result\":\"%s\"}", msg); err != nil {
		a.logger.Warn("error writing response", zap.Error(err))
	}
}

func (a HttpAdapter) respondError(w http.ResponseWriter, msg string, status int, err error) {
	a.logger.Info("error serving request", zap.Error(err))

	http.Error(w, fmt.Sprintf("{\"error\":\"%s\"}", msg), status)
}
