package wisdom

import "context"

type Client interface {
	GetRandom(ctx context.Context) (ResponseRandom, error)
}

type ResponseRandom struct {
	Text string
}
