package checker

import (
	"fmt"
	"net/http"
	"time"

	"github.com/vasilesk/word-of-wisdom/internal/service/pow/httphelper"
	"github.com/vasilesk/word-of-wisdom/pkg/logger"
	"github.com/vasilesk/word-of-wisdom/pkg/pow"
	"github.com/vasilesk/word-of-wisdom/pkg/signer"
)

type Service struct {
	l logger.Logger

	cf pow.ChallengeFactory
	s  signer.Signer

	validDuration time.Duration
}

func New(l logger.Logger, cf pow.ChallengeFactory, s signer.Signer, valid time.Duration) *Service {
	return &Service{l: l, cf: cf, s: s, validDuration: valid}
}

func (s *Service) HTTPMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodOptions {
				//nolint:contextcheck
				if err := s.submitPowInfo(w, r); err != nil {
					s.l.WithError(err).Errorf("cannot submit pow challenge info")

					w.WriteHeader(http.StatusInternalServerError)

					return
				}
			} else {
				//nolint:contextcheck
				valid, err := s.validatePowInfo(r)
				if err != nil {
					s.l.WithError(err).Errorf("cannot validate pow challenge info")

					w.WriteHeader(http.StatusUnauthorized)

					return
				}

				if !valid {
					s.l.Warnf("pow is not correct")

					w.WriteHeader(http.StatusUnauthorized)

					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}

func (s *Service) submitPowInfo(w http.ResponseWriter, r *http.Request) error {
	c, err := s.cf.GetNewChallenge(r.Context())
	if err != nil {
		return fmt.Errorf("getting challenge: %w", err)
	}

	httphelper.SubmitChallenge(w.Header(), c)

	signed, err := s.s.Sign(newPowData(
		c.String(),
		time.Now().Add(s.validDuration),
		r.Header.Get("X-Forwarded-For"),
		r.URL.RequestURI(),
	))
	if err != nil {
		return fmt.Errorf("signing data: %w", err)
	}

	httphelper.SubmitData(w.Header(), signed)

	return nil
}

func (s *Service) validatePowInfo(r *http.Request) (bool, error) {
	data, err := s.s.Restore(httphelper.FetchData(r.Header))
	if err != nil {
		return false, fmt.Errorf("restoring data: %w", err)
	}

	dataCasted, err := powDataFromMap(data.Map())
	if err != nil {
		return false, fmt.Errorf("casting data: %w", err)
	}

	if dataCasted.ValidUntil.Before(time.Now()) {
		return false, nil
	}

	if dataCasted.IP != r.Header.Get("X-Forwarded-For") {
		return false, nil
	}

	if dataCasted.URI != r.URL.RequestURI() {
		return false, nil
	}

	challenge, err := s.cf.RestoreChallenge(r.Context(), dataCasted.Challenge)
	if err != nil {
		return false, fmt.Errorf("restoring challenge: %w", err)
	}

	valid, err := challenge.Check(
		r.Context(),
		httphelper.FetchSolution(r.Header),
		httphelper.FetchData(r.Header),
	)
	if err != nil {
		return false, fmt.Errorf("checking solution: %w", err)
	}

	return valid, nil
}
