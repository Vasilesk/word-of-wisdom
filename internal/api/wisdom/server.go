package wisdom

import (
	"net/http"

	"github.com/vasilesk/word-of-wisdom/pkg/http/proto/renderer"
	"github.com/vasilesk/word-of-wisdom/pkg/http/proto/request"
	"github.com/vasilesk/word-of-wisdom/pkg/http/proto/response"
	"github.com/vasilesk/word-of-wisdom/pkg/http/server"
)

type service struct {
	r renderer.Renderer
}

func NewService(r renderer.Renderer) server.Service {
	return &service{r: r}
}

func (s *service) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/word-of-wisdom/random", s.r.Wrap(s.random))

	return mux
}

func (s *service) random(req request.Request) response.Response {
	resp := ResponseRandom{Wisdom{Text: "123"}}

	return response.Ok(resp)
}
