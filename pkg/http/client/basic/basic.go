package basic

import (
	"net/http"
	"time"

	"github.com/vasilesk/word-of-wisdom/pkg/http/client"
)

func NewClient(timeout time.Duration) client.Doer {
	return &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       timeout,
	}
}
