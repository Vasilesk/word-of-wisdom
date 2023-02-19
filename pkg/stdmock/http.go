package stdmock

import "net/http"

//go:generate mockery --with-expecter --all

type Handler interface {
	http.Handler
}

type ResponseWriter interface {
	http.ResponseWriter
}
