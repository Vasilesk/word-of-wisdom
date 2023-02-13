package checker

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	loggermock "github.com/vasilesk/word-of-wisdom/pkg/logger/mocks"
	powmock "github.com/vasilesk/word-of-wisdom/pkg/pow/mocks"
	signermock "github.com/vasilesk/word-of-wisdom/pkg/signer/mocks"
	stdmock "github.com/vasilesk/word-of-wisdom/pkg/stdmock/mocks"
)

func TestService_HTTPMiddleware(t *testing.T) {
	const (
		url           = "/my-uri"
		challengeStr  = "challenge-string"
		validDuration = time.Second
	)

	ctx := context.Background()

	method := http.MethodOptions

	now := func() time.Time {
		return time.Unix(0, 0)
	}

	l := loggermock.NewLogger(t)
	cf := powmock.NewChallengeFactory(t)
	s := signermock.NewSigner(t)

	handler := stdmock.NewHandler(t)
	writer := stdmock.NewResponseWriter(t)
	stringer := stdmock.NewStringer(t)

	req, err := http.NewRequest(method, url, nil)
	assert.NoError(t, err)

	req = req.WithContext(ctx)

	// prepare block

	clg := powmock.NewChallenge(t)
	clg.On("String").Return(challengeStr).Times(2)

	cf.On("GetNewChallenge", ctx).Return(clg, nil).Once()

	stringer.On("String").Return("").Once()

	s.On("Sign", &powData{
		Challenge:  challengeStr,
		ValidUntil: now().Add(validDuration),
		IP:         "",
		URI:        url,
	}).Return(stringer, nil)

	writer.On("Header").Return(http.Header{}).Times(2)

	writer.On("WriteHeader", http.StatusOK).Once()

	handler.On("ServeHTTP", writer, req).Once()

	// prepare block ends

	srv := New(l, cf, s, validDuration)
	srv.now = now

	handlerWithMiddleware := srv.HTTPMiddleware()(handler)
	handlerWithMiddleware.ServeHTTP(writer, req)
}

//func TestService_HTTPMiddleware(t *testing.T) {
//	type fields struct {
//		l             logger.Logger
//		cf            pow.ChallengeFactory
//		s             signer.Signer
//		validDuration time.Duration
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   func(next http.Handler) http.Handler
//	}{
//		{
//			name: "basic",
//			fields: fields{
//				l:             nil,
//				cf:            nil,
//				s:             nil,
//				validDuration: 0,
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &Service{
//				l:             tt.fields.l,
//				cf:            tt.fields.cf,
//				s:             tt.fields.s,
//				validDuration: tt.fields.validDuration,
//			}
//			if got := s.HTTPMiddleware(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("HTTPMiddleware() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
