package pow

import (
	"context"
	"fmt"

	"github.com/vasilesk/word-of-wisdom/pkg/typeutils"
)

type ChallengeFactory interface {
	GetChallenge(ctx context.Context) (Challenge, error)
}

type Challenge interface {
	Check(ctx context.Context, solution Solution, data Data) (bool, error)
	Solve(ctx context.Context, data Data) (Solution, error)
	fmt.Stringer
}

type Solution interface {
	fmt.Stringer
}

type Data interface {
	typeutils.Byter
}
