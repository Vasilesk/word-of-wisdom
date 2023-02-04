package renderer

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vasilesk/word-of-wisdom/pkg/http/proto/request"
	"github.com/vasilesk/word-of-wisdom/pkg/http/proto/response"
)

type Renderer interface {
	Wrap(h func(r request.Request) response.Response) http.Handler
}

type trackingIDGenerator interface {
	GetTrackingID(r *http.Request) string
}

type renderer struct {
	middlewares []func(next http.Handler) http.Handler
	onErrFunc   func(error)
	trackingID  trackingIDGenerator
}

func New(opts ...Option) Renderer {
	rnd := &renderer{
		middlewares: nil,
		onErrFunc:   nil,
		trackingID:  nil,
	}

	for _, opt := range opts {
		opt(rnd)
	}

	return rnd
}

func (rnd *renderer) Wrap(h func(req request.Request) response.Response) http.Handler {
	var res http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := h(request.New(w, r))

		body, err := json.Marshal(resp.Body())
		if err != nil {
			rnd.onErr(fmt.Errorf("marsalling json body: %w", err))

			return
		}

		w.Header().Set("Content-Type", "application/json")

		if trID, ok := rnd.generateTrackingID(r); ok {
			w.Header().Set("X-Tracking-Id", trID)
		}

		w.WriteHeader(resp.HTTPCode())

		_, err = w.Write(body)
		if err != nil {
			rnd.onErr(fmt.Errorf("writing json body: %w", err))

			return
		}
	})

	for i := range rnd.middlewares {
		res = rnd.middlewares[len(rnd.middlewares)-i-1](res)
	}

	return res
}

func (rnd *renderer) onErr(err error) {
	if rnd.onErrFunc != nil {
		rnd.onErrFunc(err)
	}
}

func (rnd *renderer) generateTrackingID(r *http.Request) (string, bool) {
	if rnd.trackingID == nil {
		return "", false
	}

	return rnd.trackingID.GetTrackingID(r), true
}
