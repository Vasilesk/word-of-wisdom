package static

import (
	"context"
	"math/rand"

	"github.com/vasilesk/word-of-wisdom/internal/repo/wisdomwords"
)

func NewWisdomWords() wisdomwords.Repo {
	return &repo{}
}

type repo struct{}

func (r *repo) GetRandom(_ context.Context) (wisdomwords.Wisdom, error) {
	//nolint:gosec
	ind := rand.Int() % len(storage)

	return wisdomwords.Wisdom{Text: storage[ind]}, nil
}
