package wisdom

import (
	"net/http"

	"github.com/vasilesk/word-of-wisdom/internal/repo/wisdomwords"
	"github.com/vasilesk/word-of-wisdom/pkg/http/proto/renderer"
	"github.com/vasilesk/word-of-wisdom/pkg/http/proto/request"
	"github.com/vasilesk/word-of-wisdom/pkg/http/proto/response"
	"github.com/vasilesk/word-of-wisdom/pkg/http/server"
	"github.com/vasilesk/word-of-wisdom/pkg/logger"
)

type api struct {
	l    logger.Logger
	rnd  renderer.Renderer
	repo wisdomwords.Repo
}

func NewAPI(l logger.Logger, rnd renderer.Renderer, repo wisdomwords.Repo) server.Service {
	return &api{l: l, rnd: rnd, repo: repo}
}

func (s *api) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/word-of-wisdom/random", s.rnd.Wrap(s.Random))

	return mux
}

func (s *api) Random(req request.Request) response.Response {
	resp, err := s.repo.GetRandom(req.Ctx())
	if err != nil {
		s.l.WithError(err).Errorf("got error while getting a random wisdom")

		return response.ErrServer()
	}

	return response.Ok(mapResponseRandom(resp))
}
