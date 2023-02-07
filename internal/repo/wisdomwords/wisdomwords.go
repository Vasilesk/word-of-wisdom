package wisdomwords

import "context"

type Repo interface {
	GetRandom(ctx context.Context) (Wisdom, error)
}

type Wisdom struct {
	Text string
}
