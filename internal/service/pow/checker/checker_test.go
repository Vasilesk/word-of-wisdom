//nolint:funlen
package checker

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	loggermock "github.com/vasilesk/word-of-wisdom/pkg/logger/mocks"
	powmock "github.com/vasilesk/word-of-wisdom/pkg/pow/mocks"
	signermock "github.com/vasilesk/word-of-wisdom/pkg/signer/mocks"
	stdmock "github.com/vasilesk/word-of-wisdom/pkg/stdmock/mocks"
	"github.com/vasilesk/word-of-wisdom/pkg/typeutils"
)

func TestService_HTTPMiddleware_ServeHTTP(t *testing.T) {
	t.Parallel()

	const (
		url           = "/my-uri"
		challengeStr  = "challenge-string"
		dataStr       = "data-string"
		solutionStr   = "solution-string"
		validDuration = time.Second
	)

	ctx := context.Background()

	now := func() time.Time {
		return time.Unix(0, 0)
	}

	tests := []struct {
		name                    string
		method                  string
		prepareLogger           func(l *loggermock.Logger)
		prepareChallengeFactory func(f *powmock.ChallengeFactory)
		prepareSigner           func(s *signermock.Signer)
		prepareWriter           func(w *stdmock.ResponseWriter)
	}{
		{
			name:   "options request",
			method: http.MethodOptions,
			prepareLogger: func(l *loggermock.Logger) {
				//
			},
			prepareChallengeFactory: func(f *powmock.ChallengeFactory) {
				clg := powmock.NewChallenge(t)
				clg.On("String").Return(challengeStr).Times(2)

				f.On("GetNewChallenge", ctx).Return(clg, nil).Once()
			},
			prepareSigner: func(s *signermock.Signer) {
				stringer := stdmock.NewStringer(t)
				stringer.On("String").Return("").Once()

				s.On("Sign", &powData{
					Challenge:  challengeStr,
					ValidUntil: now().Add(validDuration),
					IP:         "",
					URI:        url,
				}).Return(stringer, nil)
			},
			prepareWriter: func(w *stdmock.ResponseWriter) {
				w.On("Header").Return(http.Header{}).Times(2)
				w.On("WriteHeader", http.StatusOK).Once()
			},
		},
		{
			name:   "get request",
			method: http.MethodGet,
			prepareLogger: func(l *loggermock.Logger) {
				//
			},
			prepareChallengeFactory: func(f *powmock.ChallengeFactory) {
				clg := powmock.NewChallenge(t)
				clg.On("Check", ctx, mock.Anything, mock.Anything).Return(true, nil).Once()

				f.On("RestoreChallenge", ctx, challengeStr).Return(clg, nil).Once()
			},
			prepareSigner: func(s *signermock.Signer) {
				data := map[string]interface{}{
					"challenge":  challengeStr,
					"validUntil": float64(time.Now().Unix()) + validDuration.Seconds(),
					"ip":         "",
					"uri":        url,
				}

				s.On("Restore", mock.Anything).Return(typeutils.NewMapper(data), nil)
			},
			prepareWriter: func(w *stdmock.ResponseWriter) {
				w.On("WriteHeader", http.StatusOK).Once()
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			l := loggermock.NewLogger(t)
			tc.prepareLogger(l)

			cf := powmock.NewChallengeFactory(t)
			tc.prepareChallengeFactory(cf)

			s := signermock.NewSigner(t)
			tc.prepareSigner(s)

			w := stdmock.NewResponseWriter(t)
			tc.prepareWriter(w)

			r, err := http.NewRequest(tc.method, url, nil)
			assert.NoError(t, err)

			r = r.WithContext(ctx)

			handler := stdmock.NewHandler(t)
			handler.On("ServeHTTP", w, r).Once()

			srv := New(l, cf, s, validDuration)
			srv.now = now

			handlerWithMiddleware := srv.HTTPMiddleware()(handler)
			handlerWithMiddleware.ServeHTTP(w, r)
		})
	}
}
