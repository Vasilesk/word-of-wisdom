//nolint:funlen,maintidx
package checker

import (
	"context"
	"errors"
	"fmt"
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
		validDuration = time.Second
	)

	err := errors.New("error")

	ctx := context.Background()

	now := func() time.Time {
		return time.Date(2023, 1, 1, 0, 0, 0, 0, time.Now().Location())
	}

	tests := []struct {
		name                    string
		method                  string
		prepareLogger           func(l *loggermock.Logger)
		prepareChallengeFactory func(f *powmock.ChallengeFactory)
		prepareSigner           func(s *signermock.Signer)
		prepareWriter           func(w *stdmock.ResponseWriter)
		expectErr               bool
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
			expectErr: false,
		},
		{
			name:   "options request getting challenge error",
			method: http.MethodOptions,
			prepareLogger: func(l *loggermock.Logger) {
				l.On("WithError", fmt.Errorf("getting challenge: %w", err)).Return(l).Once()
				l.On("Errorf", "cannot submit pow challenge info").Once()
			},
			prepareChallengeFactory: func(f *powmock.ChallengeFactory) {
				f.On("GetNewChallenge", ctx).Return(nil, err).Once()
			},
			prepareSigner: func(s *signermock.Signer) {},
			prepareWriter: func(w *stdmock.ResponseWriter) {
				w.On("WriteHeader", http.StatusInternalServerError).Once()
			},
			expectErr: true,
		},
		{
			name:   "options request signing error",
			method: http.MethodOptions,
			prepareLogger: func(l *loggermock.Logger) {
				l.On("WithError", fmt.Errorf("signing data: %w", err)).Return(l).Once()
				l.On("Errorf", "cannot submit pow challenge info").Once()
			},
			prepareChallengeFactory: func(f *powmock.ChallengeFactory) {
				clg := powmock.NewChallenge(t)
				clg.On("String").Return(challengeStr).Once()

				f.On("GetNewChallenge", ctx).Return(clg, nil).Once()
			},
			prepareSigner: func(s *signermock.Signer) {
				s.On("Sign", &powData{
					Challenge:  challengeStr,
					ValidUntil: now().Add(validDuration),
					IP:         "",
					URI:        url,
				}).Return(nil, err)
			},
			prepareWriter: func(w *stdmock.ResponseWriter) {
				w.On("WriteHeader", http.StatusInternalServerError).Once()
			},
			expectErr: true,
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
					"validUntil": float64(now().Unix()) + validDuration.Seconds(),
					"ip":         "",
					"uri":        url,
				}

				s.On("Restore", mock.Anything).Return(typeutils.NewMapper(data), nil)
			},
			prepareWriter: func(w *stdmock.ResponseWriter) {
				w.On("WriteHeader", http.StatusOK).Once()
			},
			expectErr: false,
		},
		{
			name:   "get request restoring error",
			method: http.MethodGet,
			prepareLogger: func(l *loggermock.Logger) {
				l.On("WithError", fmt.Errorf("restoring data: %w", err)).Return(l).Once()
				l.On("Errorf", "cannot validate pow challenge info").Once()
			},
			prepareChallengeFactory: func(f *powmock.ChallengeFactory) {},
			prepareSigner: func(s *signermock.Signer) {
				s.On("Restore", mock.Anything).Return(nil, err)
			},
			prepareWriter: func(w *stdmock.ResponseWriter) {
				w.On("WriteHeader", http.StatusUnauthorized).Once()
			},
			expectErr: true,
		},
		{
			name:   "get request casting error challenge was not found",
			method: http.MethodGet,
			prepareLogger: func(l *loggermock.Logger) {
				l.On(
					"WithError",
					fmt.Errorf("casting data: %w", errors.New("challenge was not found")),
				).Return(l).Once()
				l.On("Errorf", "cannot validate pow challenge info").Once()
			},
			prepareChallengeFactory: func(f *powmock.ChallengeFactory) {},
			prepareSigner: func(s *signermock.Signer) {
				data := map[string]interface{}{
					"validUntil": float64(now().Unix()) + validDuration.Seconds(),
					"ip":         "",
					"uri":        url,
				}

				s.On("Restore", mock.Anything).Return(typeutils.NewMapper(data), nil)
			},
			prepareWriter: func(w *stdmock.ResponseWriter) {
				w.On("WriteHeader", http.StatusUnauthorized).Once()
			},
			expectErr: true,
		},
		{
			name:   "get request not valid until",
			method: http.MethodGet,
			prepareLogger: func(l *loggermock.Logger) {
				l.On("Warnf", "pow is not correct").Once()
			},
			prepareChallengeFactory: func(f *powmock.ChallengeFactory) {},
			prepareSigner: func(s *signermock.Signer) {
				data := map[string]interface{}{
					"challenge":  challengeStr,
					"validUntil": float64(now().Add(-time.Second).Unix()),
					"ip":         "",
					"uri":        url,
				}

				s.On("Restore", mock.Anything).Return(typeutils.NewMapper(data), nil)
			},
			prepareWriter: func(w *stdmock.ResponseWriter) {
				w.On("WriteHeader", http.StatusUnauthorized).Once()
			},
			expectErr: true,
		},
		{
			name:   "get request another ip",
			method: http.MethodGet,
			prepareLogger: func(l *loggermock.Logger) {
				l.On("Warnf", "pow is not correct").Once()
			},
			prepareChallengeFactory: func(f *powmock.ChallengeFactory) {},
			prepareSigner: func(s *signermock.Signer) {
				data := map[string]interface{}{
					"challenge":  challengeStr,
					"validUntil": float64(now().Unix()) + validDuration.Seconds(),
					"ip":         "another",
					"uri":        url,
				}

				s.On("Restore", mock.Anything).Return(typeutils.NewMapper(data), nil)
			},
			prepareWriter: func(w *stdmock.ResponseWriter) {
				w.On("WriteHeader", http.StatusUnauthorized).Once()
			},
			expectErr: true,
		},
		{
			name:   "get request another url",
			method: http.MethodGet,
			prepareLogger: func(l *loggermock.Logger) {
				l.On("Warnf", "pow is not correct").Once()
			},
			prepareChallengeFactory: func(f *powmock.ChallengeFactory) {},
			prepareSigner: func(s *signermock.Signer) {
				data := map[string]interface{}{
					"challenge":  challengeStr,
					"validUntil": float64(now().Unix()) + validDuration.Seconds(),
					"ip":         "",
					"uri":        url + "/another",
				}

				s.On("Restore", mock.Anything).Return(typeutils.NewMapper(data), nil)
			},
			prepareWriter: func(w *stdmock.ResponseWriter) {
				w.On("WriteHeader", http.StatusUnauthorized).Once()
			},
			expectErr: true,
		},
		{
			name:   "get request error restoring challenge",
			method: http.MethodGet,
			prepareLogger: func(l *loggermock.Logger) {
				l.On("WithError", fmt.Errorf("restoring challenge: %w", err)).Return(l).Once()
				l.On("Errorf", "cannot validate pow challenge info").Once()
			},
			prepareChallengeFactory: func(f *powmock.ChallengeFactory) {
				f.On("RestoreChallenge", ctx, challengeStr).Return(nil, err).Once()
			},
			prepareSigner: func(s *signermock.Signer) {
				data := map[string]interface{}{
					"challenge":  challengeStr,
					"validUntil": float64(now().Unix()) + validDuration.Seconds(),
					"ip":         "",
					"uri":        url,
				}

				s.On("Restore", mock.Anything).Return(typeutils.NewMapper(data), nil)
			},
			prepareWriter: func(w *stdmock.ResponseWriter) {
				w.On("WriteHeader", http.StatusUnauthorized).Once()
			},
			expectErr: true,
		},
		{
			name:   "get request error checking solution",
			method: http.MethodGet,
			prepareLogger: func(l *loggermock.Logger) {
				l.On("WithError", fmt.Errorf("checking solution: %w", err)).Return(l).Once()
				l.On("Errorf", "cannot validate pow challenge info").Once()
			},
			prepareChallengeFactory: func(f *powmock.ChallengeFactory) {
				clg := powmock.NewChallenge(t)
				clg.On("Check", ctx, mock.Anything, mock.Anything).Return(false, err).Once()

				f.On("RestoreChallenge", ctx, challengeStr).Return(clg, nil).Once()
			},
			prepareSigner: func(s *signermock.Signer) {
				data := map[string]interface{}{
					"challenge":  challengeStr,
					"validUntil": float64(now().Unix()) + validDuration.Seconds(),
					"ip":         "",
					"uri":        url,
				}

				s.On("Restore", mock.Anything).Return(typeutils.NewMapper(data), nil)
			},
			prepareWriter: func(w *stdmock.ResponseWriter) {
				w.On("WriteHeader", http.StatusUnauthorized).Once()
			},
			expectErr: true,
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
			if !tc.expectErr {
				handler.On("ServeHTTP", w, r).Once()
			}

			srv := New(l, cf, s, validDuration)
			srv.now = now

			handlerWithMiddleware := srv.HTTPMiddleware()(handler)
			handlerWithMiddleware.ServeHTTP(w, r)
		})
	}
}
