package renderer

import "net/http"

type Option func(rnd *renderer)

func OptionMiddleware(mm ...func(next http.Handler) http.Handler) Option {
	return func(rnd *renderer) {
		rnd.middlewares = append(rnd.middlewares, mm...)
	}
}

func OptionOnErr(onErrFunc func(err error)) Option {
	return func(rnd *renderer) {
		rnd.onErrFunc = onErrFunc
	}
}

func OptionTrackingIDGenerator(trID trackingIDGenerator) Option {
	return func(rnd *renderer) {
		rnd.trackingID = trID
	}
}
