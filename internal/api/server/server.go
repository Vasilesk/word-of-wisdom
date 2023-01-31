package server

import (
	"net/http"

	"github.com/vasilesk/word-of-wisdom/pkg/http/server"
)

type service struct {
	//
}

func NewService() server.Service {
	return &service{}
}

func (s *service) Handler() http.Handler {
	mux := http.NewServeMux()

	return mux
}
