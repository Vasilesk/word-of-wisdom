package solver

import (
	"fmt"
	"net/http"

	"github.com/vasilesk/word-of-wisdom/internal/service/pow/httphelper"
	"github.com/vasilesk/word-of-wisdom/pkg/http/client"
	"github.com/vasilesk/word-of-wisdom/pkg/pow"
	"github.com/vasilesk/word-of-wisdom/pkg/pow/gopow"
)

func NewPowSolver(c client.Doer) client.Doer {
	return &cl{c: c}
}

type cl struct {
	c client.Doer
}

func (c *cl) Do(req *http.Request) (*http.Response, error) {
	if req.Method != http.MethodOptions {
		return c.withChallengeSolving(req)
	}

	//nolint:wrapcheck
	return c.c.Do(req)
}

func (c *cl) withChallengeSolving(req *http.Request) (*http.Response, error) {
	challenge, data, err := c.obtainChallenge(req)
	if err != nil {
		return nil, fmt.Errorf("obtaining challenge: %w", err)
	}

	solution, err := challenge.Solve(req.Context(), data)
	if err != nil {
		return nil, fmt.Errorf("solving challenge: %w", err)
	}

	httphelper.SubmitSolution(req.Header, solution)
	httphelper.SubmitData(req.Header, data)

	//nolint:wrapcheck
	return c.c.Do(req)
}

//nolint:nonamedreturns
func (c *cl) obtainChallenge(initReq *http.Request) (challenge pow.Challenge, data fmt.Stringer, err error) {
	challengeReq, err := http.NewRequest(http.MethodOptions, initReq.URL.String(), nil)
	if err != nil {
		return nil, nil, fmt.Errorf("creating challenge request: %w", err)
	}

	challengeReq = challengeReq.WithContext(initReq.Context())

	resp, err := c.c.Do(challengeReq)
	if err != nil {
		return nil, nil, fmt.Errorf("doing challenge request: %w", err)
	}

	defer resp.Body.Close()

	challenge, err = gopow.NewChallenge(httphelper.FetchChallenge(resp.Header))
	if err != nil {
		return nil, nil, fmt.Errorf("constructing challenge: %w", err)
	}

	return challenge, httphelper.FetchData(resp.Header), nil
}
