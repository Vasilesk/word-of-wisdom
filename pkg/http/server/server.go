package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/vasilesk/word-of-wisdom/pkg/logger"
)

type Service interface {
	Handler() http.Handler
}

func RunServer(ctx context.Context, l logger.Logger, cfg Config, s Service) error {
	httpServer := getHTTPServer(cfg, s.Handler())

	go func() {
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
		defer cancel()

		//nolint:contextcheck
		err := httpServer.Shutdown(shutdownCtx)
		if err != nil {
			l.Errorf("shutting down server: %w", err)
		}
	}()

	err := httpServer.ListenAndServe()
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return fmt.Errorf("serving http: %w", err)
	}

	return nil
}

func getHTTPServer(cfg Config, h http.Handler) *http.Server {
	const (
		maxHeaderBytes = 1 << 20
	)

	//nolint:exhaustruct
	return &http.Server{
		Addr:           fmt.Sprintf("0.0.0.0:%d", cfg.Port),
		ReadTimeout:    cfg.RequestTimeout,
		WriteTimeout:   cfg.RequestTimeout,
		MaxHeaderBytes: maxHeaderBytes,
		Handler:        h,
	}
}
