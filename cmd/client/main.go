package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vasilesk/word-of-wisdom/internal/service/pow/solver"
	"github.com/vasilesk/word-of-wisdom/pkg/config"
	"github.com/vasilesk/word-of-wisdom/pkg/http/client"
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

	if err := exampleRequests(ctx, l, httpClient); err != nil {
		return fmt.Errorf("making example requests: %w", err)
	}

	return nil
}

func exampleRequests(ctx context.Context, l logger.Logger, httpClient client.Doer) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			if err := exampleRequest(ctx, l, httpClient); err != nil {
				return fmt.Errorf("making request: %w", err)
			}
		}
	}
}

func exampleRequest(ctx context.Context, l logger.Logger, httpClient client.Doer) error {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8090/word-of-wisdom/random", nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	req = req.WithContext(ctx)

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("making request: %w", err)
	}

	defer resp.Body.Close()

	l.WithData(map[string]interface{}{"code": resp.StatusCode}).Infof("got response")

	return nil
}
