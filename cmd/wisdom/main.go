package main

import (
	"context"
	"fmt"

	apiservice "github.com/vasilesk/word-of-wisdom/internal/api/wisdom"
	"github.com/vasilesk/word-of-wisdom/internal/repo/wisdomwords/static"
	"github.com/vasilesk/word-of-wisdom/internal/service/pow/checker"
	"github.com/vasilesk/word-of-wisdom/pkg/config"
	"github.com/vasilesk/word-of-wisdom/pkg/http/proto/renderer"
	"github.com/vasilesk/word-of-wisdom/pkg/http/server"
	"github.com/vasilesk/word-of-wisdom/pkg/logger"
	"github.com/vasilesk/word-of-wisdom/pkg/logger/zero"
	"github.com/vasilesk/word-of-wisdom/pkg/pow/gopow"
	"github.com/vasilesk/word-of-wisdom/pkg/signer/jwt"
	"github.com/vasilesk/word-of-wisdom/pkg/stopper"
)

const app = "wisdom"

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
	cfg, err := config.NewFromFile[Config]("cmd/wisdom/config/prod/config.yml")
	if err != nil {
		return fmt.Errorf("reading config: %w", err)
	}

	challengeFactory, err := gopow.NewChallengeFactory(cfg.Pow.Difficulty, cfg.Pow.NonceSize)
	if err != nil {
		return fmt.Errorf("creating challenge factory: %w", err)
	}

	signer := jwt.NewSigner(cfg.Signer.Key)

	powChecker := checker.New(l, challengeFactory, signer, cfg.PowChecker.ChallengeValid)

	rnd := renderer.New(
		renderer.OptionMiddleware(
			powChecker.HTTPMiddleware(),
		),
	)

	service := apiservice.NewAPI(l, rnd, static.NewWisdomWords())

	if err := server.RunServer(ctx, l, cfg.Server.ToServerConfig(), service); err != nil {
		return fmt.Errorf("running server: %w", err)
	}

	l.Infof("app '%s' was gracefully shut down", app)

	return nil
}
