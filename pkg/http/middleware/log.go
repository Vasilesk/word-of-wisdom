package middleware

import (
	"net/http"
	"time"

	"github.com/vasilesk/word-of-wisdom/pkg/logger"
)

func NewLogging(l logger.Logger, methods []string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t := time.Now()

			next.ServeHTTP(w, r)

			for _, method := range methods {
				if method == r.Method {
					data := map[string]interface{}{
						"duration": time.Since(t),
						"method":   r.Method,
						"uri":      r.URL.RequestURI(),
					}

					l.WithData(data).Infof("request processed")

					break
				}
			}
		})
	}
}
