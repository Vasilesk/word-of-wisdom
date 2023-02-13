//nolint:funlen
package checker

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_powDataFromMap(t *testing.T) {
	t.Parallel()

	var (
		challenge  = "challenge"
		validUntil = time.Date(2023, 1, 1, 0, 0, 0, 0, time.Now().Location())
		ip         = "ip"
		uri        = "uri"
	)

	tests := []struct {
		name        string
		m           map[string]interface{}
		expected    *powData
		expectedErr bool
	}{
		{
			name: "success",
			m: map[string]interface{}{
				"challenge":  challenge,
				"validUntil": float64(validUntil.Unix()),
				"ip":         ip,
				"uri":        uri,
			},
			expected:    newPowData(challenge, validUntil, ip, uri),
			expectedErr: false,
		},
		{
			name: "no challenge",
			m: map[string]interface{}{
				"validUntil": float64(validUntil.Unix()),
				"ip":         ip,
				"uri":        uri,
			},
			expected:    nil,
			expectedErr: true,
		},
		{
			name: "invalid challenge type",
			m: map[string]interface{}{
				"challenge":  42,
				"validUntil": float64(validUntil.Unix()),
				"ip":         ip,
				"uri":        uri,
			},
			expected:    nil,
			expectedErr: true,
		},
		{
			name: "no validUntil",
			m: map[string]interface{}{
				"challenge": challenge,
				"ip":        ip,
				"uri":       uri,
			},
			expected:    nil,
			expectedErr: true,
		},
		{
			name: "invalid validUntil type",
			m: map[string]interface{}{
				"challenge":  challenge,
				"validUntil": 42,
				"ip":         ip,
				"uri":        uri,
			},
			expected:    nil,
			expectedErr: true,
		},
		{
			name: "no ip",
			m: map[string]interface{}{
				"challenge":  challenge,
				"validUntil": float64(validUntil.Unix()),
				"uri":        uri,
			},
			expected:    nil,
			expectedErr: true,
		},
		{
			name: "invalid ip type",
			m: map[string]interface{}{
				"challenge":  challenge,
				"validUntil": float64(validUntil.Unix()),
				"ip":         42,
				"uri":        uri,
			},
			expected:    nil,
			expectedErr: true,
		},
		{
			name: "no uri",
			m: map[string]interface{}{
				"challenge":  challenge,
				"validUntil": float64(validUntil.Unix()),
				"ip":         ip,
			},
			expected:    nil,
			expectedErr: true,
		},
		{
			name: "invalid uri type",
			m: map[string]interface{}{
				"challenge":  challenge,
				"validUntil": float64(validUntil.Unix()),
				"ip":         ip,
				"uri":        42,
			},
			expected:    nil,
			expectedErr: true,
		},
	}
	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := powDataFromMap(tc.m)

			if tc.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.expected, got)
		})
	}
}
