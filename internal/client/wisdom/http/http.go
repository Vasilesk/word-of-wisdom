package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vasilesk/word-of-wisdom/internal/client/wisdom"
	"github.com/vasilesk/word-of-wisdom/pkg/http/client"
)

type cl struct {
	c       client.Doer
	baseURL string
}

func NewWisdomClient(c client.Doer, baseURL string) wisdom.Client {
	return &cl{c: c, baseURL: baseURL}
}

func (c *cl) GetRandom(ctx context.Context) (wisdom.ResponseRandom, error) {
	req, err := http.NewRequest(http.MethodGet, c.baseURL+"/random", nil)
	if err != nil {
		return wisdom.ResponseRandom{}, fmt.Errorf("creating request: %w", err)
	}

	req = req.WithContext(ctx)

	resp, err := client.DoJSON[responseRandom](c.c, req)
	if err != nil {
		return wisdom.ResponseRandom{}, fmt.Errorf("making request: %w", err)
	}

	return mapResponseRandom(resp), nil
}
