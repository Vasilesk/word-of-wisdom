package pow

import (
	"context"
	"fmt"
)

//go:generate mockery --with-expecter --name ChallengeFactory
type ChallengeFactory interface {
	GetNewChallenge(ctx context.Context) (Challenge, error)
	RestoreChallenge(_ context.Context, marshaled string) (Challenge, error)
}

//go:generate mockery --with-expecter --name Challenge
type Challenge interface {
	Check(ctx context.Context, solution Solution, data Data) (bool, error)
	Solve(ctx context.Context, data Data) (Solution, error)
	fmt.Stringer
}

type Solution interface {
	fmt.Stringer
}

type Data interface {
	fmt.Stringer
}
