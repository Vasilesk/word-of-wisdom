package main

import (
	"context"
	"fmt"

	"github.com/vasilesk/word-of-wisdom/internal/client/wisdom"
	"github.com/vasilesk/word-of-wisdom/internal/client/wisdom/wisdomhttp"
	"github.com/vasilesk/word-of-wisdom/internal/service/pow/solver"
	"github.com/vasilesk/word-of-wisdom/pkg/config"
	"github.com/vasilesk/word-of-wisdom/pkg/http/client/basic"
	"github.com/vasilesk/word-of-wisdom/pkg/logger"
	"github.com/vasilesk/word-of-wisdom/pkg/logger/zero"
	"github.com/vasilesk/word-of-wisdom/pkg/stopper"
)

const app = "powclient"

func main() {
	ctx, cancel := stopper.New(context.Background())
	defer cancel()

	l := zero.New(app)

	l.Infof("app '%s' has started", app)

	if err := run(ctx, l); err != nil {
		l.Fatalf("while running app '%s' got the next error: %v", app, err)
	}
}

func run(ctx context.Context, l logger.Logger) error {
	cfg, err := config.NewFromFile[Config]("cmd/client/config/prod/config.yml")
	if err != nil {
		return fmt.Errorf("reading config: %w", err)
	}

	httpClient := basic.NewClient(cfg.HTTPClient.Timeout)
	httpClient = solver.NewPowSolver(httpClient)

	apiClient := wisdomhttp.NewWisdomClient(httpClient, cfg.API.Wisdom.BaseURL)

	if err := exampleRequests(ctx, l, apiClient); err != nil {
		return fmt.Errorf("making example requests: %w", err)
	}

	return nil
}

func exampleRequests(ctx context.Context, l logger.Logger, c wisdom.Client) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			if err := exampleRequest(ctx, l, c); err != nil {
				return fmt.Errorf("making request: %w", err)
			}
		}
	}
}

func exampleRequest(ctx context.Context, l logger.Logger, c wisdom.Client) error {
	resp, err := c.GetRandom(ctx)
	if err != nil {
		return fmt.Errorf("getting random from client: %w", err)
	}

	l.WithData(map[string]interface{}{"text": resp.Text}).Infof("got response")

	return nil
}
