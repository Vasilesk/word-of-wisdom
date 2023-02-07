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

func FetchSolution(r *http.Request) fmt.Stringer {
	return typeutils.NewStringer(r.Header.Get(headerSolution))
}

func FetchData(r *http.Request) fmt.Stringer {
	return typeutils.NewStringer(r.Header.Get(headerData))
}

func SubmitData(w http.ResponseWriter, data fmt.Stringer) {
	w.Header().Add(headerData, data.String())
}

func SubmitChallenge(w http.ResponseWriter, challenge fmt.Stringer) {
	w.Header().Add(headerChallenge, challenge.String())
}
