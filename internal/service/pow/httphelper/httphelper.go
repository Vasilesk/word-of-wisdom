package httphelper

import (
	"fmt"
	"net/http"

	"github.com/vasilesk/word-of-wisdom/pkg/typeutils"
)

const (
	headerChallenge = "X-Pow-Challenge"
	headerData      = "X-Pow-Data"

	headerSolution = "X-Pow-Solution"
)

func FetchChallenge(h http.Header) string {
	return h.Get(headerChallenge)
}

func SubmitChallenge(h http.Header, challenge fmt.Stringer) {
	h.Add(headerChallenge, challenge.String())
}

func FetchData(h http.Header) fmt.Stringer {
	return typeutils.NewStringer(h.Get(headerData))
}

func SubmitData(h http.Header, data fmt.Stringer) {
	h.Add(headerData, data.String())
}

func FetchSolution(h http.Header) fmt.Stringer {
	return typeutils.NewStringer(h.Get(headerSolution))
}

func SubmitSolution(h http.Header, data fmt.Stringer) {
	h.Add(headerSolution, data.String())
}
