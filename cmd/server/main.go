package main

import (
	"context"
	"fmt"

	apiservice "github.com/vasilesk/word-of-wisdom/internal/api/server"
	"github.com/vasilesk/word-of-wisdom/pkg/config"
	"github.com/vasilesk/word-of-wisdom/pkg/http/proto/renderer"
	"github.com/vasilesk/word-of-wisdom/pkg/http/server"
	"github.com/vasilesk/word-of-wisdom/pkg/logger"
	"github.com/vasilesk/word-of-wisdom/pkg/logger/zero"
	"github.com/vasilesk/word-of-wisdom/pkg/stopper"
)

const app = "server"

func main() {
	ctx, cancel := stopper.New(context.Background())
	defer cancel()

	l := zero.New(app)

	if err := run(ctx, l); err != nil {
		l.Fatalf("while running app '%s' got the next error: %v", app, err)
	}
}

func run(ctx context.Context, l logger.Logger) error {
	l.Infof("app '%s' has started", app)

	cfg, err := config.NewFromFile[Config]("cmd/server/config.yml")
	if err != nil {
		return fmt.Errorf("reading config: %w", err)
	}

	rnd := renderer.New()

	service := apiservice.NewService(rnd)

	if err := server.RunServer(ctx, l, cfg.Server.ToServerConfig(), service); err != nil {
		return fmt.Errorf("running server: %w", err)
	}

	l.Infof("app '%s' was gracefully shut down", app)

	return nil
}
