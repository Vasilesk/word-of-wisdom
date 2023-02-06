package gopow

import (
	"context"
	"errors"
	"fmt"

	gopow "github.com/bwesterb/go-pow"

	"github.com/vasilesk/word-of-wisdom/pkg/pow"
	"github.com/vasilesk/word-of-wisdom/pkg/typeutils"
)

type challenge struct {
	challenge          gopow.Request
	challengeMarshaled string
}

func NewChallenge(difficulty int, nonce []byte) (pow.Challenge, error) {
	if difficulty <= 0 {
		return nil, errors.New("difficulty should be positive")
	}

	challengeMarshaled := gopow.NewRequest(uint32(difficulty), nonce)

	var c gopow.Request

	if err := c.UnmarshalText([]byte(challengeMarshaled)); err != nil {
		return nil, fmt.Errorf("unmarshalling challenge: %w", err)
	}

	return &challenge{
		challenge:          c,
		challengeMarshaled: challengeMarshaled,
	}, nil
}

func (c *challenge) Check(_ context.Context, solution pow.Solution, data pow.Data) (bool, error) {
	res, err := gopow.Check(c.challengeMarshaled, solution.String(), data.Bytes())
	if err != nil {
		return false, fmt.Errorf("checking solution: %w", err)
	}

	return res, nil
}

func (c *challenge) Solve(_ context.Context, data pow.Data) (pow.Solution, error) {
	proof, err := c.challenge.Fulfil(data.Bytes()).MarshalText()
	if err != nil {
		return nil, fmt.Errorf("marshalling proof: %w", err)
	}

	return typeutils.NewStringer(string(proof)), nil
}

func (c *challenge) String() string {
	// no error can be returned by the library
	// no need to change the interface because of it
	res, _ := c.challenge.MarshalText()

	return string(res)
}
