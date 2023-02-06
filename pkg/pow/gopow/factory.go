package gopow

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"

	"github.com/vasilesk/word-of-wisdom/pkg/pow"
)

func NewChallengeFactory(difficulty int, nonceSize int) (pow.ChallengeFactory, error) {
	if difficulty <= 0 {
		return nil, errors.New("difficulty should be positive")
	}

	if nonceSize <= 0 {
		return nil, errors.New("nonceSize should be positive")
	}

	return &factory{
		difficulty: difficulty,
		nonceSize:  nonceSize,
	}, nil
}

type factory struct {
	difficulty int
	nonceSize  int
}

func (f *factory) GetNewChallenge(_ context.Context) (pow.Challenge, error) {
	nonce := make([]byte, f.nonceSize)

	if _, err := rand.Read(nonce); err != nil {
		return nil, fmt.Errorf("generating random nonce: %w", err)
	}

	return NewRandomChallenge(f.difficulty, nonce)
}

func (f *factory) RestoreChallenge(_ context.Context, marshaled string) (pow.Challenge, error) {
	return NewChallenge(marshaled)
}
