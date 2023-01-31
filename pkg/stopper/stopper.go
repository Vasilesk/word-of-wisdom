package stopper

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func New(parent context.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(parent)

	signals := make(chan os.Signal, 1)

	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signals
		cancel()
	}()

	return ctx, cancel
}
