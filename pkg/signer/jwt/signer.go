package jwt

import (
	"errors"
	"fmt"

	jwt "github.com/golang-jwt/jwt/v4"

	"github.com/vasilesk/word-of-wisdom/pkg/signer"
	"github.com/vasilesk/word-of-wisdom/pkg/typeutils"
)

type signerJWT struct {
	key []byte
}

func NewSigner(key string) signer.Signer {
	return &signerJWT{key: []byte(key)}
}

func (s *signerJWT) Sign(data signer.Data) (signer.Signed, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("got error casting the claims")
	}

	for k, v := range data.Map() {
		claims[k] = v
	}

	tokenStr, err := token.SignedString(s.key)
	if err != nil {
		return nil, fmt.Errorf("while signing: %w", err)
	}

	return typeutils.NewStringer(tokenStr), nil
}

func (s *signerJWT) Restore(signed signer.Signed) (signer.Data, error) {
	token, err := jwt.Parse(signed.String(), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("got error casting signed value")
		}

		return s.key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("parsing signed jwt: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("token is invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("claims are not a map")
	}

	return typeutils.NewMapper[string, interface{}](claims), nil
}
